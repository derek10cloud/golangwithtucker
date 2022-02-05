package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} //Declare variable and Initialize variable

	for i := 0; i < 5; i++ { //loop for array "t"
		fmt.Printf("%v ", t[i]) //print arrat "t"'s elements
	}
	fmt.Println()

	var nums [5]int
	// nums[0] -> 0
	// nums[1] -> 0
	// ...
	fmt.Println(nums)
	days := [3]string{"mon", "tue", "wed"}
	fmt.Println(days)
	var temps [5]float64 = [5]float64{24.3, 26.7}
	fmt.Println(temps)

	var s = [5]int{1: 10, 3: 30}
	fmt.Println(s) // first element is 10, third element is  30

	x := [...]int{10, 20, 30}
	fmt.Println(x)
}
