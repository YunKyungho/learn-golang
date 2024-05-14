package main

import "fmt"

func main() {
	defineArray()
	defineSlice()
}

func defineArray() {
	arr := [4]int{1, 2, 3}
	arr[3] = 4
	// arr[4] = 5 -> 오류 발생
	// 위 처럼 array 정의 시 길이가 4로 고정된다.
	fmt.Println(arr)
}

func defineSlice() {
	ex_slice := []int{1, 2}
	// ex_slice[2] = 3 -> 오류 발생
	// slice는 요소를 추가할 때 append 함수를 사용한다.
	ex_slice = append(ex_slice, 3)
	// append는 새로운 요소가 추가 된 slice를 반환하기에 값을 꼭 받아야 한다.
	fmt.Println(ex_slice)
}
