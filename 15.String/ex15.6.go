package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str))     //string 타입 길이
	fmt.Printf("len(runes) = %d\n", len(runes)) //[]rune 타입 길이
}
