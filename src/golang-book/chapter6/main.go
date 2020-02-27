package main

import "fmt"

/* Array Practice

func main() {
	x := [5]float64{100, 99, 98, 97, 96}
	total := 0.0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
*/

// Slice Practice
/*
func main() {
	x := make([]int, 5)
	for i := range x {
		x[i] = i
	}
	for i := range x {
		x = append(x, i)
	}
	x2 := make([]int, len(x))
	copy(x2, x)
	fmt.Println(x)
	fmt.Println(x2)
}
*/

// Map Practice
/*
func main() {
	newMap := make(map[string]string)
	newMap["Deez"] = "Nutz"
	fmt.Println(newMap)
	value, ok := newMap["Deez"]
	if ok {
		fmt.Println("We Found", value)
	}
	delete(newMap, "Deez")
	if value, ok := newMap["Deez"]; ok {
		fmt.Println("We Found", value)
	} else {
		fmt.Println("Ligma Balls")
	}
	fmt.Println(newMap)
}
*/

// Question 4
func main() {
	arr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := 1000000
	for _, value := range arr {
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println(smallest)
}
