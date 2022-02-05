package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) //매개변수로 입력받은 파일 열기
	if err != nil {
		return "", err
	} //err가 nil이면 err 내용 출력
	defer file.Close()             //함수가 끝나면 열었던 파일을 닫음
	rd := bufio.NewReader(file)    //bufio패키지의 NewReader 객체 생성
	line, _ := rd.ReadString('\n') //ReadString 메서드로 \n가 나올때까지 읽음 -> 한줄을 읽어서 line 변수에 저장.
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) //매개변수로 입력받은 filename으로 파일을 생성
	if err != nil {
		return err
	} //err가 nil이면 err 내용 출력
	defer file.Close()                //함수가 끝나면 열었던 파일 닫음
	_, err = fmt.Fprintln(file, line) //파일에 line이라는 문자열 작성
	//Fprintln은 첫번째 반환 값으로 쓴 길이를 반환하고, 두번째로는 error발생 시 error를 반환함
	return err
}

const filename string = "data.txt" //filename 변수 선언

func main() {
	line, err := ReadFile(filename) //파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") //파일 읽으려 했는데, err발생하면, WriteFile 함수 호출
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		} //WriteFile함수에서도 err가 발생(return하는 err가 nil이 아니면)하면, 파일 생성 실패했다고 출력
		line, err = ReadFile(filename) //WriteFile을 통해서 다시 한 번 file읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.")
			return
		} // 그래도 err가 발생하면 파일 읽기 실패
	}
	fmt.Println("파일내용:", line) //ReadFile함수로 파일을 읽을 수 있으면, 읽은 파일의 첫번째 줄 출력
}
