package main

import "fmt"

// 내장함수 이용할 수도 있음 func copy(dst, src []Type) int

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10) //len: 3, cap:10
	slice3 := make([]int, 10)    //len: 10, cap:10

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	//slice2 := make([]int, len(slice1))
	//copy (slice2, slice1)
	//copy 이용한 것은 2줄이나 써야되서 append가 더 나은듯

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)

	//요소 삭제하기
	slice1 = append(slice1[:2], slice1[3:]...) //요소 삭제하는 쉬운 방법, ...을 붙여줘여 동작하는 듯
	fmt.Println(slice1)

	//요소 추가하기 1 - 삭제된 3 추가
	slice1 = append(slice1[:2], append([]int{3}, slice1[2:]...)...) //요소 추가 방법 1
	//slice1 = append(slice1[:2], 3, slice1[2:]... )
	fmt.Println(slice1)

	//요소 추가하기 2 - 3 뒤에 100 추가 (불필요한 메모리 사용을 막으려고...?)
	slice1 = append(slice1, 0)   //맨 뒤에 요소 추가 -> [1 2 3 4 5 0]
	copy(slice1[4:], slice1[3:]) // [4 5 0] --복사--> [5 0]  => slice1은 [1 2 3 4 4 5]
	slice1[3] = 100
	fmt.Println(slice1)
	// 복잡함...
}
