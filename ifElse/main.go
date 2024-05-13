package main

import "fmt"

func main() {
	fmt.Println(commonIf(28))
	fmt.Println(defineVariableIf(28))
	fmt.Println(useSwitch(39))
}

func commonIf(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func defineVariableIf(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func useSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
