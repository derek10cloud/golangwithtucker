package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1바이트
	C int8 //1바이트
	E int8 //1바이트
	B int  //8바이트
	D int  //8바이트
} // 총 24바이트
//메모리 패딩은 임베디드 같은 메모리에 제한이 있는 프로그래밍할 때 필요함

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
