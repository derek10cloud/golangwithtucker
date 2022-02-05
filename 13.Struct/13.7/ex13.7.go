package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1바이트
	B int  //8바이트
	C int8 //1바이트
	D int  //8바이트
	E int8 //1바이트
} //실제로는 19바이트 크기
//메모리 패딩 때문에 1바이트 짜리 모두 8바이트에 맞추기 위해 7바이트씩 더해서 8바이트씩 할당받고 총 40바이트로 메모리 패딩 됨

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
