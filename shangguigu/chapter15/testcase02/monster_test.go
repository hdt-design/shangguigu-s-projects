package monster

import (
	"testing"
)

func TestStore(t *testing.T) {

	//创建monster实例
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "三味真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误,希望为%v,实际为%v",
			true, res)
	}
	t.Logf("monster.Store()测试成功")
}

func TestRestore(t *testing.T) {
	var monster = &Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.ReStore()错误,希望为%v,实际为%v",
			true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误,希望为%v,实际为%v",
			"红孩儿", monster.Name)
	}
	t.Logf("monster.Restore()测试成功")
}
