// Double array

package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, //you have to make comma even though, there is no another array.
	} //중괄호가 바로 위에 줄과 같이 있었으면 쉼표를 안찍는다.
	for i, arr := range a {
		for i, v := range arr {
			fmt.Print(v, " ", i)
		}
		fmt.Println()
	}
}
