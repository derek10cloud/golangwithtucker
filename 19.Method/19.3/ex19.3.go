package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)  //포인터 메서드 호출
	fmt.Println(mainA.balance) // -30 되어서 70 출력

	mainA.withdrawValue(20)    //값 메서드 호출
	fmt.Println(mainA.balance) // 값만 복사 되었으므로 그대로 70 출력

	var mainB account = mainA.withdrawReturnValue(20) //값 메서드를 호출하였지만 return이 있으므로 -20되어서 50 리턴 -> 값복사 2번 일어남
	fmt.Println(mainB.balance, mainB)                 //50 호출

	mainB.withdrawPointer(30) //포인터 메서드 호출
	fmt.Println(mainB.balance) // -30 되어서 20 출력
}
