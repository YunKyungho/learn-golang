package main

import (
	"fmt"
	"github.com/YunKyungho/learn-golang/basic/method/myStruct"
)

func main() {
	account := myStruct.NewAccount("ygh")
	account.Deposit(10)
	fmt.Println(account.Balance())
	fmt.Println(account)
}
