package main

import (
	"fmt"
)

// 定义一个结构体Account

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法
// 1.存款

func (account *Account) Deposite(money float64, pwd string) {

	//验证密码
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	//验证存款金额
	if money <= 0 {
		fmt.Println("存款金额必须大于0")
		return
	}

	//存款
	account.Balance += money
	fmt.Println("存款成功")
}

// 2.取款

func (account *Account) Withdraw(money float64, pwd string) {

	//验证密码
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	//验证取款金额
	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额有误")
		return
	}

	//取款
	account.Balance -= money
	fmt.Println("取款成功")
}

// 3.查询余额

func (account *Account) QueryBalance(pwd string) {

	//验证密码
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)
}

func main() {

	//测试
	account := Account{
		AccountNo: "gs1111111",
		Pwd:       "123456",
		Balance:   1000,
	}

	//这里可以做的更加灵活，就是让用户通过控制台来输入命令...
	//菜单....
	account.QueryBalance("123456")
	account.Deposite(200.0, "123456")
	account.QueryBalance("123456")
	account.Withdraw(150.0, "123456")
	account.QueryBalance("123456")

}
