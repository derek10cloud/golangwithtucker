// slice 동작원리
// type SliceHeader struct{
// 	Data uintptr //실제 배열을 가리키는 포인터 -> 배열을 가르키는 포인터 덕분에 쉽게 크기가 다른 배열을 가르키도록 변경할 수 있다.
// 	Len int //요소 개수
// 	Cap int //실제 배열의 길이 (용량) -> 용량은 배열 전체 길이 에서 시작 인덱스를 뺀 값
// }

package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}

//지난주에 일반적으로 함수에 들어가는 인수는 값으로 취급되기 때문에 일반 배열은 값이 복사, but slice는 메모리 주소가 복사
//값이 복사되지 않고, 슬라이스의 필드인 "포인터", "길이", "용량" 만 복사 됨

// func changeArray2(array2 *[5]int) {
// 	array2[2] = 200
// }
