package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a //p1은 int형 값을 가지는 a변수의 주소값
	var p2 *int = &a //p2는 int형 값을 가지는 a변수의 주소값
	var p3 *int = &b //p3는 int형 값을 가지는 b변수의 주소값

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}
