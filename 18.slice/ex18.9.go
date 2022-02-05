// 두 슬라이스가 같은 배열을 가르켜서 생기는 문제들을 해결할 수 있는 방법
// 슬라이스 복제
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	// 이전에는 slice2 := append(slice1, 4, 5) 방식으로 slice1이 가리키는 배열에 변화를 준 것
	// 지금은 slice1이랑 길이만 같은 새로운 배열을 만들어서 값을 복사하는 것

	for i, v := range slice1 {
		slice2[i] = v
	}

	// 그런데 이렇게 순회를 이용해서 요솟값 복사하는 것은 복잡
	// slice2 := append([]int{}, slice1) 으로도 복사할 수 있음
	// 위 방법은 slice2 := append([]int{}, slice1[0], slice1[1], ...)와 같음
	// 내장함수 이용할 수도 있음 func copy(dst, src []Type) int -> 다음예제

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
