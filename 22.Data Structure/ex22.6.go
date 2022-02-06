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

	for k, v := range m {
		fmt.Println(k, v)
	}
}
