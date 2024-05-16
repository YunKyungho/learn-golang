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
}
