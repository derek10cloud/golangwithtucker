package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User { //User타입의 주소 타입을 반환
	var u = User{name, age}
	return &u //탈출 분석으로 u메모리가 사라지지 않음
}

//원래 함수가 종료되면서 u변수는 사용되지 않으므로
//u의 주소값도 사라지는 것이 맞는데, 왜 주소가 잘 출력이 되었을까?
//함수 내의 지역변수는 스택 메모리 영역(함수 끝나면 사라지는)에 쌓이는데,
//Go의 Escape Analysis 덕분에, 컴파일 시 코드를 분석하고 인스턴스가 바깥으로 탈출한다는 것이 분석되면,
//힙 메모리 영역에 넣어서, 함수에서 탈출해도 사라지지 않게함
//힙 메모리 영역에 저장된 값은 쓰임이 있으면 사라지지 않음.
//여기서는 var userPointer에서 u의 주소값이 저장되었으므로, &u는 쓰임이 있고 사라지지 않는다.
//Go에서는 컴파일 시 탈출 분석을 통해서 스택에 만들지 힙에 만들지 결정함

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
