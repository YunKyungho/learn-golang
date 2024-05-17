package main

import "fmt"

func main() {
	loopWithRange()
	loopWithoutRange()
	result := loopWithParameter(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

func loopWithRange() {
	var arr = [3]int{1, 2, 3}
	for index, element := range arr {
		fmt.Println(index, element)
	}
}

func loopWithoutRange() {
	var arr = [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}

func loopWithParameter(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
