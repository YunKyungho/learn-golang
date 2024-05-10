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
}

// 마찬가지로 main package를 찾고 그 내부의 main 함수를 찾아 실행시킨다.
// 물론 컴파일이 아닌 package의 공유를 목적으로 코드를 작성한다면 굳이 main이 있을 필요는 없다.

func runPrint() {
	fmt.Println("Hello world!")
	// func를 export 하려면 대문자로 작성해야한다.
	// mymod.go 파일 참고
	mymod.SayHello()
	mymod.sayBye()
	// 소문자로 시작하는 함수는 export 불가능하다.
}
