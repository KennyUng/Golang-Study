package main

import "fmt"

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x *int, y *int) {
	hold := *x
	*x = *y
	*y = hold
}
