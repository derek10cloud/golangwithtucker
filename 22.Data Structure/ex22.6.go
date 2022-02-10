package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)
	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[328] = Product{"샤프", 3000}
	m[5432] = Product{"샤프심", 500}

	fmt.Printf("%s은 %v 원 입니다.\n", m[16].Name, m[16].Price)
	fmt.Printf("%s은 %v 원 입니다.\n", m[16].Name, m[46].Price)

	for k, v := range m { //배열에서 인덱스와 값이 저장되듯, 키와 값이 저장됨
		fmt.Println(k, v) //출력을 보면 정렬되어 있지 않음
		//맵은 내부에서 요소를 보관할 때 입력 순서나 키 값과  상관없이 보관함 
	}
}
