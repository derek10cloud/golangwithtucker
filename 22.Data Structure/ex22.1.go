package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()      //리스트 생성
	v.PushBack(4)        //리스트 맨 뒤에 요소 추가
	e1 := v.PushFront(1) //리스트 맨 앞에 요소 추가
	v.InsertAfter(2, e1) //e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() { //각 요소 순회
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { //각 요소 역 순회
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
