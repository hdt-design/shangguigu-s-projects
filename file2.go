// Package file 提供文件管理相关的功能，包括上传、删除、恢复、
package file

import (
	"fmt"      // 导入fmt包，用于格式化IO操作
	"net/http" // 导入net/http包，用于HTTP服务器和客户端功能
	"os"       // 导入os包，用于操作系统级的交互
	"path/filepath" // 导入path/filepath包，用于处理文件路径
	"strconv"  // 导入strconv包，用于字符串和数字之间的转换
	"time"     // 导入time包，用于处理时间相关操作

	"github.com/gin-gonic/gin" // 导入gin框架，用于构建Web服务器
	"gorm.io/gorm"            // 导入gorm包，用于ORM操作
)

// File 定义文件结构体，包含文件的相关信息
type File struct {
	ID        uint64    `gorm:"primarykey"` // 文件ID，主键
	UserID    uint64    `gorm:"index"`        // 用户ID，索引
	Filename  string    `gorm:"size:255"`   // 文件名，最大长度255
	Filepath  string    `gorm:"size:512"`   // 文件路径，最大长度512
	Filesize  int64                        // 文件大小
	SharedKey string                         // 共享密钥
	DeletedAt *time.Time                   // 删除时间，如果文件未被删除则为nil
	CreatedAt time.Time                      // 创建时间
}

// favoriteFile 定义收藏文件结构体，包含用户收藏的文件信息
type favoriteFile struct {
	ID     uint64 `gorm:"primarykey"` // 收藏记录ID，主键
	UserID uint64 `gorm:"index"`        // 用户ID，索引
	FileID uint64 `gorm:"index"`        // 文件ID，索引
}

// 声明一个全局变量db，类型为*gorm.DB，用于存储数据库连接
var db *gorm.DB

// InitDB 初始化数据库连接，接收一个*gorm.DB类型的参数database
func InitDB(database *gorm.DB) {
    // 将传入的database赋值给全局变量db
	db = database
    // 根据File和favoriteFile结构体自动迁移数据库表结构
	db.AutoMigrate(&File{}, &favoriteFile{})
}

// UploadFile 处理文件上传请求，接收一个*gin.Context类型的参数c
func UploadFile(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 从请求表单中获取名为"file"的文件，并检查是否有错误发生
	file, err := c.FormFile("file")
	if err != nil {
        // 如果有错误发生，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
    // 根据用户ID创建上传路径，路径格式为"./uploads/userID/"
	uploadPath := fmt.Sprintf("./uploads/%d/", userID)
    // 确保上传路径存在，如果不存在则创建该路径，并设置权限为最大权限
	os.MkdirAll(uploadPath, os.ModePerm)
    // 将文件路径和文件名拼接，生成文件的完整路径
	filePath := filepath.Join(uploadPath, file.Filename)
    // 将上传的文件保存到指定路径，并检查是否有错误发生
	if err := c.SaveUploadedFile(file, filePath); err != nil {
        // 如果有错误发生，返回状态码500（Internal Server Error）和错误信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
    // 创建一个新的File结构体实例，填充文件信息
	newFile := File{
		UserID:   userID,
		Filename: file.Filename,
		Filepath: filePath,
		Filesize: file.Size,
	}
    // 将新文件记录插入到数据库中
	db.Create(&newFile)
    // 返回状态码200（OK）和上传成功的消息以及文件ID
	c.JSON(http.StatusOK, gin.H{
		"message": "upload success", "file_id": newFile.ID,
	})
}

// DeleteFile 删除用户的文件，接收一个*gin.Context类型的参数c
func DeleteFile(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 从请求URL查询参数中获取file_id
	fileID := c.Query("file_id")
    // 声明一个File类型的变量file，用于存储查询到的文件信息
	var file File
    // 根据file_id和user_id查询文件，如果文件不存在则返回错误信息
	if err := db.Where("id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
    // 尝试从文件系统中删除文件，如果删除失败则返回错误信息
	if err := os.Remove(file.Filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from storage"})
		return
	}
    // 获取当前时间
	now := time.Now()
    // 更新数据库中文件的deleted_at字段为当前时间，表示文件已被删除
	db.Model(&file).Update("deleted_at", &now)
    // 返回状态码200（OK）和删除成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// RestoreFile 恢复用户的文件，接收一个*gin.Context类型的参数c
func RestoreFile(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 从请求URL路径参数中获取file_id
	fileID := c.Param("file_id")
    // 声明一个File类型的变量file，用于存储查询到的文件信息
	var file File
    // 根据file_id和user_id查询文件，如果文件不存在则返回错误信息
	if err := db.Where("id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
    // 更新数据库中文件的deleted_at字段为nil，表示文件已被恢复
	db.Model(&file).Update("deleted_at", nil)
    // 返回状态码200（OK）和恢复成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "File restored successfully"})
}

// FavoriteFile 收藏用户的文件，接收一个*gin.Context类型的参数c
func FavoriteFile(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 从请求URL查询参数中获取file_id
	fileIDStr := c.Query("file_id")
    // 将file_id字符串转换为uint64类型，如果转换失败则返回错误信息
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file_id"})
		return
	}
    // 创建一个新的favoriteFile结构体实例，填充收藏信息，并插入到数据库中
	db.Create(&favoriteFile{
		UserID: userID,
		FileID: fileID,
	})
    // 返回状态码200（OK）和收藏成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "File favorited successfully"})
}

// stringToUint64 将字符串转换为uint64类型，接收一个string类型的参数s
func stringToUint64(s string) uint64 {
    // 声明一个uint64类型的变量result，用于存储转换后的结果
	var result uint64
    // 使用fmt.Sscanf将字符串s按照"%d"格式扫描并转换为uint64类型，存储在result中
	fmt.Sscanf(s, "%d", &result)
    // 返回转换后的uint64类型结果
	return result
}

// UnfavoriteFile 取消收藏用户的文件，接收一个*gin.Context类型的参数c
func UnfavoriteFile(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 从请求URL查询参数中获取file_id
	fileIDStr := c.Query("file_id")
    // 将file_id字符串转换为uint64类型，如果转换失败则返回错误信息
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file_id"})
		return
	}
    // 根据user_id和file_id从数据库中删除收藏记录
	db.Where("user_id = ? AND file_id = ?", userID, fileID).Delete(&favoriteFile{})
    // 返回状态码200（OK）和取消收藏成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "File unfavorited successfully"})
}

// ListFavoriteFiles 列出用户收藏的文件，接收一个*gin.Context类型的参数c
func ListFavoriteFiles(c *gin.Context) {
    // 从gin.Context中获取user_id，并将其转换为uint64类型
	userID := c.MustGet("user_id").(uint64)
    // 声明一个favoriteFile类型的切片favFiles，用于存储查询到的收藏记录
	var favFiles []favoriteFile
    // 根据user_id查询所有收藏记录，并存储在favFiles中
	db.Where("user_id = ?", userID).Find(&favFiles)
    // 声明一个uint64类型的切片fileIDs，用于存储收藏的文件ID
	var fileIDs []uint64
    // 遍历收藏记录，将每个记录的FileID添加到fileIDs切片中
	for _, fav := range favFiles {
		fileIDs = append(fileIDs, fav.FileID)
	}
    // 声明一个File类型的切片files，用于存储查询到的文件信息
	var files []File
    // 根据fileIDs切片中的文件ID查询所有文件信息，并存储在files切片中
	db.Where("id IN ?", fileIDs).Find(&files)
    // 返回状态码200（OK）和收藏的文件列表
	c.JSON(http.StatusOK, gin.H{"favorite_files": files})
}
