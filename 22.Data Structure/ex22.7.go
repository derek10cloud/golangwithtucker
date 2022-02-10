package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3) //요솟값 삭제 시도
	delete(m, 4) //없는 요솟값 삭제 시도
	fmt.Println(m[3]) //삭제된 요솟값 출력
	fmt.Println(m[1]) //존재하는 요솟값 출력
	fmt.Println(m[5]) //아예 만들어지지도 않은 요솟값 출력
}
