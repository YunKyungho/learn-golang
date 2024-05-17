package main

import "fmt"

func main() {
	defineMap()
}

func defineMap() {
	ex_map := map[string]int{"height": 173, "weight": 75}
	ex_map["age"] = 28
	// 키, 값 생성
	// ex_map[29] = "map" -> 오류 발생
	// 지정된 타입으로만 요소 추가가 가능하다.
	loopWithMap(ex_map)
	checkKey(ex_map)
}

func loopWithMap(iter map[string]int) {
	for key, value := range iter {
		// 필요 없는 값은 _(언더바)를 기재한다.
		fmt.Println(key, value)
		fmt.Println(iter[key])
		// 키로 값 조회
	}
}

func checkKey(iter map[string]int) {
	val, exists := iter["dfe"]
	// go에서는 key로 map을 조회 시 2개의 값을 return 한다.
	// 첫번째는 키에 상응하는 값, 두번째는 키가 존재하는지 안 하는지에 대한 bool 값.
	if !exists {
		println("Not in age")
	}
	println(val)
}
