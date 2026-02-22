package monster

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给Monster绑定方法Store，可以将Monster变量序列化后保存到文件中

func (m *Monster) Store() bool {

	//序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return false
	}

	//保存到文件
	filePath := "/Users/liuxinpei/Desktop/monster.json"
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (m *Monster) Restore() bool {
	filePath := "/Users/liuxinpei/Desktop/monster.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
