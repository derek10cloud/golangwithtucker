package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler { //1.핸들러 인스턴스 생성 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   //2. Student 객체를 []byte로 변환
	w.Header().Add("content-type", "application/json") //3. JSON 포맷 명시
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) //4. 결과 전송
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
