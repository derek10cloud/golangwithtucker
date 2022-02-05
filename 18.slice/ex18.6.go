package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [1 2 3], 길이: 3, 용량: 3
	// slice2 배열: [1 2 3 4 5 빔], 길이: 5, 용량: 6 => slice1 배열과 다른 배열을 가르키고 있다.
	// 기존 배열에 append 진행 시 기존 배열의 용량이 부족할 때는, 더 큰 배열을 마련해서 기존 배열 요소를 복사한다.
	// 용량은 부족할 때 2배로

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [1 100 3], 길이: 3, 용량: 3
	// slice2 배열: [1 2 3 4 5 빔], 길이: 5, 용량: 6 => slice1 배열과 다른 배열을 가르키고 있다.

	slice1 = append(slice1, 500)
	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	// slice1 배열: [1 100 3 500 빔 빔], 길이: 4, 용량: 6
	// slice2 배열: [1 2 3 4 5 빔], 길이: 5, 용량: 6 => slice1 배열과 다른 배열을 가르키고 있다.
}
