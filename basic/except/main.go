package main

import (
	"fmt"
	"github.com/YunKyungho/learn-golang/basic/except/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("ygh")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
