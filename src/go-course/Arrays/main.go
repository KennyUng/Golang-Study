package main

import "fmt"

func main() {
	// Array Calls
	/*
		var x [5]int
		fmt.Println(x)
		x[3] = 42
		fmt.Println(x)
		x2 := [...]int{1, 2, 3, 4, 5}
		fmt.Println(x2)
	*/

	// Slice Calls

	// x := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5}
	x := []int{1, 3, 2, 3, 4, 3, 5, 3, 6, 3}
	target := 3

	for i := 0; i < len(x); i++ {
		if x[i] == target {
			x = append(x[:i], x[i+1:]...)
			i--
		}
	}
	fmt.Println(x)
}
