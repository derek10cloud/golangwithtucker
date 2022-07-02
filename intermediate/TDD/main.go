package main

import (
	"fmt"
)

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)

	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
	opMap["**"] = pow
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func pow(a, b int) int {
	rst := 1
	for i := 0; i < b; i++ {
		rst *= a
	}
	return rst
}

func Calculate(op string, a, b int) int {
	// if op == "+" {
	// 	return a + b
	// } else if op == "-" {
	// 	return a - b
	// } else if op == "*" {
	// 	return a * b
	// }

	// rst := 0
	// switch op {
	// case "+":
	// 	rst = a + b
	// case "-":
	// 	rst = a - b
	// case "*":
	// 	rst = a * b
	// case "/":
	// 	rst = a / b
	// }
	// return rst

	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func main() {
	initOpMap()
	test()
}

func test() {
	// o := Calculate("+", 3, 2)
	// if o != 5 {
	// 	fmt.Printf("Test1 Failed! expected:%d output:%d\n", 5, o)
	// 	return
	// }
	// fmt.Println("Success!")

	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	// o = Calculate("+", 5, 4)
	// if o != 9 {
	// 	fmt.Printf("Test2 Failed! expected:%d output:%d\n", 9, o)
	// 	return
	// }
	// fmt.Println("Success!")

	if !testCalculate("Test1", "+", 5, 4, 9) {
		return
	}

	if !testCalculate("Test3", "-", 5, 3, 2) {
		return
	}

	if !testCalculate("Test4", "-", 3, 6, -3) {
		return
	}

	if !testCalculate("Test5", "*", 3, 7, 21) {
		return
	}

	if !testCalculate("Test6", "*", 3, 0, 0) {
		return
	}

	if !testCalculate("Test7", "*", 3, -3, -9) {
		return
	}

	if !testCalculate("Test8", "/", 9, 3, 3) {
		return
	}

	if !testCalculate("Test9", "**", 2, 3, 8) {
		return
	}

	if !testCalculate("Test10", "**", 3, 3, 27) {
		return
	}

	if !testCalculate("Test11", "**", 2, 0, 1) {
		return
	}

	fmt.Println("Success!")
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}
