package mymod

import "fmt"

// 소문자로 작성된 func은 private func이 되어 외부에서 호출할 수 없다.
func sayBye() {
	fmt.Println("Bye")
}

// 외부로 export를 위해선 대문자로 시작하는 func을 작성해야 한다.
func SayHello() {
	fmt.Println("Hello")
}
