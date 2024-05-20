package main

import (
	"fmt"
	"time"
)

func main() {
	go task1()
	task2()
	// go task2() -> main이 바로 종료.
	// main 함수는 goroutine을 기다려주지 않는다.
	// task1의 4가 출력되기 전에 프로그램이 종료되는 경우도 간혹 있다. 이 또한 위의 이유 때문이다.
}

func task1() {
	arr := [4]int{1, 2, 3, 4}
	for _, elm := range arr {
		time.Sleep(time.Second)
		fmt.Println(elm)
	}
}

func task2() {
	arr := [4]int{5, 6, 7, 8}
	for _, elm := range arr {
		time.Sleep(time.Second)
		fmt.Println(elm)
	}
}
