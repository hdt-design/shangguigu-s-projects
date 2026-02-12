// Package model
package model

import (
	"fmt"
)

// 定义一个结构体Account
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数--构造函数

func Newaccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度有误")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度有误")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额不能小于20元")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

// 方法
// 1.存款
func (account *account) Deposit(money float64, pwd string) {

	//验证密码
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return
	}

	//验证存款金额
	if money <= 0 {
		fmt.Println("存款金额有误")
		return
	}

	//存款
	account.balance += money
	fmt.Println("存款成功，当前余额为：", account.balance)
}

// 2.取款
func (account *account) Withdraw(money float64, pwd string) {

	//验证密码
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return
	}

	//验证取款金额
	if money <= 0 {
		fmt.Println("取款金额有误")
		return
	}

	//取款
	if account.balance < money {
		fmt.Println("余额不足")
		return
	}

	account.balance -= money
	fmt.Println("取款成功，当前余额为：", account.balance)
}

// 3.查询余额
func (account *account) Balance(pwd string) float64 {

	//验证密码
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return 0
	}

	return account.balance
}
