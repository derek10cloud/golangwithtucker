package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
		//파일 출력 Fprintf
	})

	http.ListenAndServe(":3000", nil)
}
