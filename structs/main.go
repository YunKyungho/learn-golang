package main

import "fmt"

func main() {
	menuPrice := map[string]int{"kimbab": 4000, "tunaKimbab": 6000}
	kimbabHeaven := restaurant{
		name:           "kimbabHeaven",
		classification: "snackBar",
		menuPrice:      menuPrice,
	}
	// python의 keyword 인자 처럼 사용이 가능하다.
	fmt.Println(kimbabHeaven)
	fmt.Println(kimbabHeaven.name)
}

type restaurant struct {
	name           string
	classification string
	menuPrice      map[string]int
}

// struct 내부에 method 생성도 가능하나 생성자 함수는 없다는 것을 기억해야한다.
// 생성자가 필요하다면 우리가 직접 method를 실행 시켜줘야만 한다.
