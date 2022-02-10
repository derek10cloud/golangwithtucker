package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) //5개의 요소를 가진 링 자료 구조 생성, r은 첫번째 요소 가리키는 포인터

	n := r.Len() //링 길이 

	for i := 0; i < n; i++ {
		r.Value = 'A' + i //순회하면서 모든 요소에 값 대입
		//r이 0일 때 A, 1:B, 2:C, 3:D, 4:E
		r = r.Next()      // r이 4일 때 요소 값은 E인데, r에 그 다음 요소(r=0)를 저장했음
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value) //순회하면서 출력
		r = r.Next()
	}
	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Prev()
	}
	fmt.Println()
}
