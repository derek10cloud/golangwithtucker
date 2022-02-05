//slicing

package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5}
	slice := array[1:2] //새로운 배열을 만드는 것이 아닌 배열의 일부를 포인터로 가리키는 슬라이스를 생성할 뿐

	fmt.Println("array:", array, len(array), cap(array))
	fmt.Println("slice:", slice, len(slice), cap(slice))
	//slice는 array를 가리키고 있고 인덱스 1부터 시작하므로 시작 인덱스는 1 cap은 배열의 총길이에서 시작 인덱스 뺀 값. [1 빔 빔 빔 빔], 길이:1, 용량:4

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)
	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}
