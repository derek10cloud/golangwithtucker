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

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
	fmt.Println()
}
