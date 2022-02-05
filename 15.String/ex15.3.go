package main

import "fmt"

func main() {
	var char rune = '😁'

	fmt.Printf("%T\n", char) //rune 타입은 int32와 같아서 int32출력
	fmt.Println(char)        //char값 출력, int32라 숫자로 출력
	fmt.Printf("%c\n", char) //%c포맷 사용하면 문자 출력
}
