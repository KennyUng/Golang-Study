package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	catchphrase string
}

type secretAgent struct {
	person
	ltk bool
}

func (p *person) speak() string {
	return p.catchphrase
}

func main() {
	jb := secretAgent{
		person: person{
			firstName:   "James",
			lastName:    "Bond",
			catchphrase: "Martini. Shaken, not stirred.",
		},
		ltk: true,
	}

	mm := person{
		firstName:   "Miss",
		lastName:    "Moneypenny",
		catchphrase: "Hello James.",
	}

	fmt.Println(jb.firstName, jb.lastName, "says", jb.speak())
	fmt.Println(mm.firstName, mm.lastName, "says", mm.speak())
}
