package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List //List 구조체를 필드로 가지는 구조체 정의
}

func (q *Queue) Push(val interface{}) { //요소를 추가하는 메서드
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} { //요소 반환하면서 삭제하는 메서드
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue { //새로운 리스트를 만들어서 Queue 구조체에 저장하는 함수
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue() //새로운 큐 생성

	for i := 1; i < 5; i++ { //요소 뒤에서 입력하는 메서드 호출
		queue.Push(i)
	}
	v := queue.Pop() //요소 하나 반환하고 삭제하는 메서드 호출 (1번째)
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop() //요소 하나 반환하고 삭제하는 메서드 호출 순회 (2번째~마지막)
	}
	fmt.Println()
}
