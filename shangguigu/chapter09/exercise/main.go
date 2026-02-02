package main

import "fmt"

/*
1)使用map[string]map[string]string的map类型
2）key：表示用户名
3）如果某个用户存在，将密码修改为“888888”
4）如果某个用户不存在，则添加该用户
5）编写一个函数modifyuser(users map[string]map[stringstring,name string)

*/

func modifyuser(users map[string]map[string]string, name string) {

	//判断users是否存在该用户
	//v,ok:=users[name]
	if users[name] != nil {
		//有该用户，修改密码
		users[name]["pwd"] = "888888"
	} else {
		//没有该用户，添加该用户
		users[name] = make(map[string]string, 2)
		users[name]["name"] = name
		users[name]["pwd"] = "888888"
	}
}
func main() {
	uesrs := make(map[string]map[string]string, 10)
	uesrs["smith"] = make(map[string]string, 2)
	uesrs["smith"]["name"] = "smi"
	uesrs["smith"]["pwd"] = "123456"

	modifyuser(uesrs, "tom")
	modifyuser(uesrs, "jerry")
	modifyuser(uesrs, "smith")

	fmt.Println(uesrs)
}
