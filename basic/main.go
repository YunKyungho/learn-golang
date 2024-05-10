package main

// Go의 컴파일러는 패키지 이름이 main인 것 부터 찾기에 컴파일을 위해 main package가 필수적이다.

import (
	"fmt"

	"github.com/YunKyungho/learn-golang/basic/mymod"
	// 위 같은 형식으로 import 하려면 프로젝트 root 디렉토리에서 아래 명령어를 실행한다.
	// go mod init github.com/YunKyungho/learn-golang
)

func main() {
	runPrint()
	defineVarAndConst()
}

// 마찬가지로 main package를 찾고 그 내부의 main 함수를 찾아 실행시킨다.
// 물론 컴파일이 아닌 package의 공유를 목적으로 코드를 작성한다면 굳이 main이 있을 필요는 없다.

func runPrint() {
	fmt.Println("Hello world!")
	// func를 export 하려면 대문자로 작성해야한다.
	// mymod.go 파일 참고
	mymod.SayHello()
	// mymod.sayBye()
	// 소문자로 시작하는 함수는 export 불가능하다.
}

func defineVarAndConst() {
	const name string = "Kyungho"
	fmt.Println(name)
	// 상수 선언 방법
	// go는 정적타입 언어로 사용할 type을 지정해야한다.
	// 당연하게도 상수는 선언한 뒤 name = "Gyeongho" 처럼 값의 변경이 불가하다.

	var job string = "Developer"
	fmt.Println(job)
	// 기본적인 변수 선언 방법

	age := 28
	fmt.Println(age)
	// 축약형 변수 선언 방법
	// func 내부에서만 가능한 문법으로 위 처럼 간단하게 변수 선언도 가능하며 이를 '타입 추론'이라고 한다.
	// '저게 되면 동적 타입 언어가 아니냐?' 할 수 있지만 컴파일 시점에 컴파일러가 타입을 결정하는 것이기에
	// 실행 시점에 변수의 타입이 결정되는 python 같은 동적 타입 언어와는 엄연히 다르다.
	// age = "28"
	// 그 예로 golang에서는 위 처럼 변수의 타입을 바꾸는 코드 실행이 불가능하다.

}
