package main

import "fmt"

func main() {
	var feet float32
	fmt.Scanf("%f", &feet)
	meters := (feet * .3048)
	fmt.Println(meters)
}
