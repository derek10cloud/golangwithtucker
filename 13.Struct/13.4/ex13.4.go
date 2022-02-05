//필드 중복 (상귀 구조체의 필드명이 서로 겹칠 때)
package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("유저: %s ID: %s 나이: %d VIP레벨: %d 유저레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,      //Level이라는 필드가 같을 경우, 현재 변수타입에 선언된 Level에 접근한다. 즉 VIP 구조체의 Level을 의미함
		vip.User.Level, //상위 구조체의 필드를 호출하고 싶으면, 상위 구조체의 타입까지 지정해줘야 함.
	)
}
