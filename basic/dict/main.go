package main

import (
	"fmt"
	"github.com/YunKyungho/learn-golang/basic/dict/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
