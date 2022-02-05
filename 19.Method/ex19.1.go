package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}  //a는 account{balance:100}인 인스턴스의 주소 값을 a에 할당
	withdrawFunc(a, 30) //함수 형태 호출
	//포인터 변수를 할당하는 이유는 이 함수에 주소 값을 전달하려고! 계속 봤지만 함수에 주솟값을 안넣고 일반 값을 넣으면... 매개변수랑 일반변수랑 서로 다른 값을 갖게 되므로...!

	a.withdrawMethod(30) //메서드 형태 호출

	fmt.Printf("%d \n", a.balance)
}
