package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [0 0 0 빔 빔], 길이: 3, 용량: 5
	// slice2 배열: [0 0 0 4 5], 길이: 5, 용량: 5
	// 만일 slice1이 용량이 4짜리였다면, 빈 값에 5를 넣을 때 용량을 초과하면서 다른 주솟값에다가 새로운 더 큰 배열을 할당함(일반적으로 2배 크기로)

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [0 100 0 빔 빔] 길이: 3, 용량: 5
	// slice2 배열: [0 100 0 4 5] 길이: 5, 용량: 5

	slice1 = append(slice1, 500)

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [0 100 0 500 빔] 길이: 4, 용량: 5
	// slice2 배열: [0 100 0 500 5] 길이: 5, 용량: 5
}
