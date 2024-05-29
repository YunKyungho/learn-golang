package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	arr := [5][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
	}

	for _, elm := range arr {
		go sumArray(c, elm)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(<-c)
	}
}

func sumArray(c chan int, arr [3]int) {
	sum := 0
	for _, elm := range arr {
		time.Sleep(time.Second)
		sum += elm
	}
	c <- sum
}
