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
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}
