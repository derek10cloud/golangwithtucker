package main

import "fmt"

//struct 타입명의 첫 번째 글자가 대문자면 패키지 외부로 공개되는 타입입니다.
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House //초기화
	house.Address = "서울특별시 서초구 교보타워"
	house.Size = 28
	house.Price = 20.832
	house.Type = "아파트"

	//초기화 방법 2
	// var house House = House{"서울특별시 서초구 교보타워", 28, 20.832, "아파트"}
	// 각각 1:1 매칭되어서 값이 들어간다.

	//일부 필드 초기화
	// var house House = House{Size:28,
	//	Type:"아파트",
	//}
	// 나머지 값은 0이나 "" 로 초기화 된다.

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d\n", house.Size)
	fmt.Printf("가격: %.2f 억원\n", house.Price)
	fmt.Println("타입:", house.Type)
}
