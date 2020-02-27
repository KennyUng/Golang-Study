package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type Rectangle struct {
	length float64
	width  float64
}

type Person struct {
	name string
}

type Android struct {
	Person
	model string
}

func (person *Person) talk() {
	fmt.Println("Hello my name is", person.name)
}

func (rec *Rectangle) area() float64 {
	return rec.length * rec.width
}

func (rec *Rectangle) perimeter() float64 {
	return 2*rec.length + 2*rec.width
}

func main() {
	/*
		c := new(Circle)
		c.x = 0.0
		c.y = 0.0
		c.r = 5.0

		c2 := Circle{0, 0, 5}
		fmt.Println(*c)
		fmt.Println(c2)
		fmt.Println(c2.area())

		rec := Rectangle{8, 10}
		fmt.Println("Area:", rec.area(), "Perimeter:", rec.perimeter())
	*/

	android := new(Android)
	android.name = "Kevin"
	android.talk()
}
