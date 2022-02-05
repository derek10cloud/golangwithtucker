package main

import "fmt"

type Data struct { //Data 타입 구조체
	value int
	data  [200]int //200개의 int타입의 요소를 갖는 배열
}

func ChangeData(arg Data) { //매개변수(파라미터)로 Data를 받음
	arg.value = 999
	arg.data[100] = 999 //data 배열의 100번째 요소의 값을 999로 대입
}

func main() {
	var data Data    //1608바이트 짜리 Data 타입 구조체 선언
	ChangeData(data) //인수(아규먼트)로 data 변수 입력
	//함수의 인자로 쓰이면 무조건 값으로 쓰임(우변)
	// a(우변, 공간) = 10 (좌변, 값)
	// data의 값이 arg라는 변수에 복사된다.
	// arg의 주소와 data의 주소가 다르므로 arg.value를 바꿔도, data의 value는 바뀌지 않음

	fmt.Printf("value = %d\n", data.value)         //data의 필드 출력
	fmt.Printf("data[100] = %d\n", data.data[100]) //data의 필드 출력
}
