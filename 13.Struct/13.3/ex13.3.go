// embeded structure

package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 필드명 생략
	VIPLevel int
	Price    int
}

func main() {
	user := User{"김진짜", "jinzza", 23}
	vip := VIPUser{
		User{"호날두", "ronaldo", 38},
		3,
		250}
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d \n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
	)
}
