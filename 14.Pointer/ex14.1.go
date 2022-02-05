package main

import "fmt"

func main() {
	var a int = 500
	var p *int //int 포인터 변수 p선언

	//a의 메모리 주소를 변수 p의 값으로 대입

	if p == nil {
		fmt.Println("Y")
	}

	fmt.Printf("p의 값: %p\n", p)            //메모리 주소값 출력
	fmt.Printf("p가 가르키는 메모리의 값: %d\n", *p) //p가 가르키는 메모리 값 출력

	*p = 100                   //p거 거루카눈 메모리 공간(a)의 값을 변경
	fmt.Printf("a의 값:%d\n", a) //a값 변화ㅖ
}
