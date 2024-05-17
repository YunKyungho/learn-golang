package myStruct

import "fmt"

type Account struct {
	owner   string
	balance int
}

// NewAccount go에서의 constructor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit go에서의 method
func (a *Account) Deposit(amount int) {
	// 여기서 a는 receiver라고 불리며 a의 type은 Account이다.
	// struct의 첫 글자를 소문자로 사용하는 것이 관행이다.
	// receiver를 작성할 때 *을 붙히지 않는다면 go에서는 생성했던 Account의 복사본을 receiver로 넘겨준다.
	// 값의 변경 사항이 복사본에 적용되고 원본 Account에는 아무런 영향을 끼치지 않는 것이다.
	// 그렇기에 꼭 method의 receiver를 작성할 때는 *을 붙혀야한다.
	a.balance += amount

}

// Balance go에서의 getter
func (a *Account) Balance() int {
	return a.balance
}

// Owner go에서의 getter
func (a *Account) Owner() string {
	return a.owner
}

// String go에서의 magic method
func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
	// fmt.Println(account) 실행 시 &{owner, balance} 형식으로 출력됬던 값이
	// ygh's account.
	// Has: 10
	// 위 처럼 출력되는 것을 볼 수 있다.
	// Println 함수 사용 시 전달 받은 struct의 String method를 자동으로 호출하는 듯 하다.
}
