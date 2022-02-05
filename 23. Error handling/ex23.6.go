package main

import "fmt"

func f() {
	fmt.Println("f()함수 시작") 
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r) //8. recover 실행 되었으므로 패닉 복구 시도
		}
	}()
	g() //2
	fmt.Println("f()함수 끝") 
} //7. 패닉 f()에 전파되면서 f()종료하려고 하는데 종료전에 defer func() 실행

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3)) //3. h() 실행
	fmt.Printf("9 / 0 = %d\n", h(9, 0)) //4. h() 실행 //6. 패닉 g()에 전파되면서 g() 종료
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.") //5. 패닉 발생 //9. 패닉 복구가 되고 패닉 메세지 출력
	}
	return a / b 
}

func main() {
	f() // 1. f() 실행
	fmt.Println("프로그램이 계속 실행됨") //10. f() 끝난 다음문장 실행
}
