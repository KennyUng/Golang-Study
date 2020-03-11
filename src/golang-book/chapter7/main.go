package main

import "fmt"

/*
func main() {
	slice1 := []float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, value := range slice1 {
		total += value
	}
	fmt.Println(average(total, float64(len(slice1))))
}

func average(total float64, len float64) float64 {
	return (total / len)
}
*/

// Better implementation
/*
func main() {
	slice := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(slice))
}

func average(slice []float64) float64 {
	total := 0.0
	for _, value := range slice {
		total += value
	}
	return total / float64(len(slice))
}
*/

// Multiple return values
/*
func main() {
	x, y := 5, 4
	sum, diff := f(x, y)
	fmt.Println("Sum:", sum, "Diff:", diff)
}

func f(x int, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}
*/

// Variadic Function Calls
/*
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
func sum(args ...int) int {
	total := 0
	for _, values := range args {
		total += values
	}
	return total
}
*/

// Deferrence
/*
func main() {
	first := func() {
		fmt.Println("First")
	}
	second := func() {
		fmt.Println("Second")
	}
	defer second()
	first()
}
*/

// Panic and Recover
/*
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Deez Nutz")
}
*/

// Question 1
/*
func main() {
	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(sum(1, 2))
}
*/

//Question 2
/*
func main() {
	fmt.Println(half(3))
}
func half(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
*/

// Question 5
func main() {
	fibVal := fib(25)
	fmt.Println(fibVal)
}
func fib(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}
