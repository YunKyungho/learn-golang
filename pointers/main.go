package main

import "fmt"

func main() {
	valueSubstitution()
	addressSubstitution()
}

func valueSubstitution() {
	a := 5
	b := a
	a = 12
	fmt.Println(a, b)
	// 12, 5가 출력 된다
}

func addressSubstitution() {
	a := 5
	b := &a
	a = 12
	fmt.Println(a, *b)
	// 12, 12가 출력된다.
	otherWay()
}

func otherWay() {
	a := 5
	b := &a
	*b = 20
	fmt.Println(a, *b)
	// 20, 20이 출력되며 b로도 a의 값을 변경할 수 있다.
}
