package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
} //PasswordError 구조체의 메서드로 Error()를 선언했기 때문에
//PasswordError 구조체는 error 인터페이스로 사용될 수 있음

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	} //암호 길이가 짧을 때 PasswordError 구조체 정보를 반환함
	// return 타입이 error이지만, PasswordError는 error 인터페이스로 쓸 수 있으므로
	// PasswordError 구조체를 반환할 수 있음
	return nil //password의 길이가 8이상이면 nil 반환
}

func main() {
	err := RegisterAccount("myID", "myPw") //ID, PW 입력
	if err != nil {                        //error가 발생하면
		if errInfo, ok := err.(PasswordError); ok { //PasswordError 메서드 호출, golang에서 자주 사용되는 구문
			fmt.Printf("%v Len: %d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		} //errInfo로 RegisterAccount의 각 필드에도 접근 가능
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
