package main

import "fmt"

type Data struct { //Data 타입 구조체
	value int
	data  [200]int //200개의 int타입의 요소를 갖는 배열
}

func ChangeData(arg *Data) { //매개변수(파라미터)로 Data타입을 사용하는 변수의 주소값을 받음
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	ChangeData(&data) //인수(아규먼트)로 data 변수의 주소값 입력
	//함수의 인자로 쓰이면 무조건 값으로 쓰임(우변)
	// 이지만, data변수의 주소값을 입력
	// data의 주소값이 arg라는 변수에 복사된다.

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
