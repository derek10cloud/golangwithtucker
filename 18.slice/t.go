package main

import "fmt"

func main() {
	var slice = make([]int, 3, 5)  // [0 0 0] len:3 cap:5
	slice = append(slice, 4, 5, 6) // [0 0 0 4 5 6] len:6 cap:10
	fmt.Println(slice, len(slice), cap(slice))
}
