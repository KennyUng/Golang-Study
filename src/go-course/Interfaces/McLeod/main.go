package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.firstName, p.lastName)
}

func main() {
	joe := person{
		"Joe",
		"Schmoe",
	}

	joe.speak()
}
