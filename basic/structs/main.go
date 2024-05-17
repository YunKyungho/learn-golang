package main

import (
	"fmt"
	"github.com/YunKyungho/learn-golang/basic/structs/public"
)

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

	ps := public.PublicStruct{AnyString: "any", AnyInt: 93}
	fmt.Println(ps)
	// public.privateStruct -> 오류 발생

	ps.AnyString = "diff"
	// 다만 public으로 내부 변수를 정의할 경우 위 처럼 아무렇게나 값을 바꿔버릴 수가 있다.
	// 또한 로직을 통해 생성된 값만 변수에 담고 싶을 수도 있다. (다른 개발자가 임의로 값을 변경하지 못하게)

	useConstructorPs := public.PublicFunc("any")
	fmt.Println(*useConstructorPs)
	// useConstructorPs.anyString -> 오류 발생
	// useConstructorPs.anyInt -> 오류 발생
	// 위 방식으로 생성자를 사용하고 내부 변수에 접근하지 못하게 만든다.
	// 근데 이렇게 private 하게 변수를 만들어서 안정성을 높였다고 했을 때 내부 변수의 값을 변경하려면 어떻게 해야 하는가?
	// ->> method를 만들어서 활용해야 한다.
}

type restaurant struct {
	name           string
	classification string
	menuPrice      map[string]int
}

// struct 내부에 method 생성도 가능하나 생성자 함수는 없다는 것을 기억해야한다.
// 생성자가 필요하다면 우리가 직접 method를 실행 시켜줘야만 한다.
