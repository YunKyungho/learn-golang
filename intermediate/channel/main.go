package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// channel 생성 방법, make 함수를 사용하며 전달할 정보의 타입을 지정해야한다.
	go task1(c)
	// goroutine 실행 시 채널을 함수의 인자로 전달
	go task2(c)
	result1 := <-c
	// 작업의 결과를 받는 방법.
	// 작업의 결과를 기다리는 것과도 같은 의미다.
	fmt.Println(result1)
	result2 := <-c
	fmt.Println(result2)
}

func task1(c chan int) {
	sum := 0
	arr := [4]int{1, 2, 3, 4}
	for _, elm := range arr {
		time.Sleep(time.Second)
		sum += elm
	}
	c <- sum
	// 전달 받은 채널로 결과를 전달
}

func task2(c chan int) {
	sum := 0
	arr := [4]int{5, 6, 7, 8}
	for _, elm := range arr {
		time.Sleep(time.Second)
		sum += elm
	}
	c <- sum
}
