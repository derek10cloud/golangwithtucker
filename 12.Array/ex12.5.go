package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
	fmt.Println()

	b = a //var a and var b's type is same, so a's value can be copied. and array's length have to be same if you want copy array's vaule
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
