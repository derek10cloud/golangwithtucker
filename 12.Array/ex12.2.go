package main

import "fmt"

const Y int = 3

func main() {
	x := 5
	//a := [x]int{1, 2, 3, 4, 5} //if you want to declare array variable's number, you can't use variable
	b := [Y]int{1, 2, 3}
	var c [10]int

	fmt.Println(x, b, c)
}
