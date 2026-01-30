package file

import (
	"net/http"       // 导入net/http包，用于HTTP服务器和客户端功能
	"strings"        // 导入strings包，用于字符串操作
	"time"           // 导入time包，用于处理时间相关操作

	"github.com/gin-gonic/gin"   // 导入gin框架，用于构建Web服务器
	"github.com/golang-jwt/jwt/v5" // 导入jwt包，用于JWT认证
	"golang.org/x/crypto/bcrypt" // 导入bcrypt包，用于密码哈希
	"gorm.io/driver/mysql"       // 导入gorm/mysql包，用于MySQL数据库驱动
	"gorm.io/gorm"               // 导入gorm包，用于ORM操作
)

// DB 返回全局数据库连接对象
func DB() *gorm.DB {
	return db
}

// AuthMiddleware 是一个中间件，用于检查Authorization头中的JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        // 获取Authorization头信息
		tokenString := c.GetHeader("Authorization")
        // 如果Authorization头为空，返回状态码401（Unauthorized）和错误信息
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

        // 去掉Authorization头中的"Bearer "前缀
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // 解析JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
        // 如果解析失败，返回状态码401（Unauthorized）和错误信息
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

        // 获取JWT的声明部分
		claims := token.Claims.(jwt.MapClaims)
        // 将声明中的username设置到gin.Context中
		c.Set("username", claims["username"].(string))
        // 将声明中的user_id转换为uint64类型并设置到gin.Context中
		c.Set("user_id", uint64(claims["user_id"].(float64)))
        // 继续处理请求
		c.Next()
	}
}

// User 模型，定义用户结构体
type User struct {
	ID       uint64 `gorm:"primarykey"` // 用户ID，主键
	Username string `gorm:"unique;size:255"` // 用户名，唯一且最大长度255
	Password string `gorm:"size:255"` // 密码，最大长度255
	Avatar   string `gorm:"size:255"` // 头像，最大长度255
	Bio      string `gorm:"size:512"` // 个人简介，最大长度512
}

// 声明一个全局变量db，类型为*gorm.DB，用于存储数据库连接
var db *gorm.DB
// 声明一个全局变量jwtKey，类型为[]byte，用于JWT签名
var jwtKey = []byte("secret")

// InitDB 初始化数据库连接
func InitDB() {
    // 声明一个error变量，用于存储错误信息
	var err error
    // 打开MySQL数据库连接，并传入连接字符串和配置
	db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/pan?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
    // 如果连接失败，抛出panic
	if err != nil {
		panic("failed to connect database")
	}
    // 自动迁移User表结构
	db.AutoMigrate(&User{})
}

// Register 注册新用户，接收一个*gin.Context类型的参数c
func Register(c *gin.Context) {
    // 声明一个匿名结构体，用于绑定请求中的JSON数据
	var user struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
	}
    // 将请求体中的JSON数据绑定到user结构体中，如果绑定失败则返回状态码400（Bad Request）和错误信息
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // 检查用户名和密码是否为空，如果为空则返回状态码400（Bad Request）和错误信息
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"})
		return
	}

    // 对用户密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    // 如果哈希处理失败，返回状态码500（Internal Server Error）和错误信息
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

    // 创建一个新的User结构体实例，填充用户信息
	newUser := User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

    // 将新用户记录插入到数据库中，如果插入失败则返回状态码500（Internal Server Error）和错误信息
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

    // 返回状态码200（OK）和注册成功的消息以及用户ID和用户名
	c.JSON(http.StatusOK, gin.H{
		"message":  "User registered successfully",
		"user_id":  newUser.ID,
		"username": newUser.Username,
	})
}

// Login 登录用户并返回JWT token，接收一个*gin.Context类型的参数c
func Login(c *gin.Context) {
    // 声明一个匿名结构体，用于绑定请求中的JSON数据
	var user struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
	}
    // 将请求体中的JSON数据绑定到user结构体中，如果绑定失败则返回状态码400（Bad Request）和错误信息
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // 声明一个User类型的变量dbUser，用于存储查询到的用户信息
	var dbUser User
    // 根据用户名查询用户，如果用户不存在则返回状态码401（Unauthorized）和错误信息
	if err := db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

    // 对输入的密码与数据库中的密码哈希进行比较，如果不匹配则返回状态码401（Unauthorized）和错误信息
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

    // 创建一个新的JWT token，包含用户ID、用户名和过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

    // 生成签名后的JWT token字符串，如果生成失败则返回状态码500（Internal Server Error）和错误信息
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

    // 返回状态码200（OK）和登录成功的消息以及生成的JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

// GetUserInfo 根据用户ID获取用户信息，接收一个*gin.Context类型的参数c
func GetUserInfo(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	uid := c.MustGet("user_id").(uint64)
    // 声明一个User类型的变量user，用于存储查询到的用户信息
	var user User
    // 根据用户ID查询用户信息，如果用户不存在则返回状态码404（Not Found）和错误信息
	if err := db.First(&user, uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

    // 返回状态码200（OK）和用户信息
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateUserInfo 更新用户的头像和简介信息，接收一个*gin.Context类型的参数c
func UpdateUserInfo(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	uid := c.MustGet("user_id").(uint64)
    // 声明一个匿名结构体，用于绑定请求中的JSON数据
	var userUpdates struct {
		Avator string `json:"avator"` // 头像
		Bio    string `json:"bio"`    // 个人简介
	}
    // 将请求体中的JSON数据绑定到userUpdates结构体中，如果绑定失败则返回状态码400（Bad Request）和错误信息
	if err := c.ShouldBindJSON(&userUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // 声明一个User类型的变量user，用于存储查询到的用户信息
	var user User
    // 根据用户ID查询用户信息，如果用户不存在则返回状态码404（Not Found）和错误信息
	if err := db.First(&user, uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

    // 如果用户更新的头像信息不为空，则更新数据库中的头像信息
	if userUpdates.Avator != "" {
		db.Model(&user).Update("avatar", userUpdates.Avator)
	}
    // 如果用户更新的简介信息不为空，则更新数据库中的简介信息
	if userUpdates.Bio != "" {
		db.Model(&user).Update("bio", userUpdates.Bio)
	}

    // 返回状态码200（OK）和更新成功的消息
	c.JSON(http.StatusOK, gin.H{
		"message": "User info updated successfully",
	})
}

// 启动gin服务器
func main() {
    // 初始化数据库连接
	InitDB()
    // 创建一个默认的gin路由引擎
	r := gin.Default()

    // 注册处理用户注册请求的路由
	r.POST("/register", Register)
    // 注册处理用户登录请求的路由
	r.POST("/login", Login)

    // 创建一个路由组，用于处理需要认证的用户信息请求
	auth := r.Group("/user")
    // 使用AuthMiddleware中间件保护该路由组
	auth.Use(AuthMiddleware())
    {
        // 注册处理获取用户信息请求的路由
		auth.GET("/info", GetUserInfo)
        // 注册处理更新用户信息请求的路由
		auth.PUT("/update", UpdateUserInfo)
	}

    // 启动gin服务器，监听8080端口
	r.Run(":8080")
}
