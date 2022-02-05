package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 난 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}

	// 인터페이스 사용하지 않고 메서드 사용도 가능함
	// fmt.Printf("%s\n", student.String())

	var stringer Stringer
	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
