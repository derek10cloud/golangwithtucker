package main

import "fmt"

type myInt int //1. 사용자 정의 별칭 타입

//2. myInt 별칭 타입을 리시버로 갖는 메서드
func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10       //myInt 타입 변수
	fmt.Println(a.add(30)) //3. myInt 타입의 add() 메서드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50)) //4. int 타입을 타입 변화
}
