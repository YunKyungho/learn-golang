package public

type PublicStruct struct {
	AnyString string
	AnyInt    int
	// struct 명 뿐만 아니라 내부의 변수 또한 대문자, 소문자로 public, private이 나뉜다.
}

type privateStruct struct {
	anyString string
	anyInt    int
	// 외부에서 호출 불가능하다.
}

func PublicFunc(anyString string) *privateStruct {
	ps := privateStruct{anyString: anyString, anyInt: 0}
	return &ps
	// 포인터를 return 하는 이유는 복사본을 반환하는 것 보다 효율적이기 때문이다.
}
