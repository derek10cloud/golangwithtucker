package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   //4바이트
	Score float64 //8바이트
}

//합치면 12바이트지만, 12바이트는 8의 배수가 아니기 때문에 메모리 정렬을 위해, 필드 사이에 공간을 띄우는 메모리 패딩을 한다.
//float64가 8바이트여서, User 구조체는 최소 8바이트 배수로 메모리 패딩이 되야 함

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user)) //그래서 실제 사이즈는 16바이트
}
