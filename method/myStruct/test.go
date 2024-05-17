package myStruct

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

func (a *Account) Balance() int {
	return a.balance
}
