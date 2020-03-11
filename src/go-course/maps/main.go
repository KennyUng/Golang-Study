package main

import "fmt"

func main() {
	newMap := make(map[string]int)
	newMap["Deez Nutz"] = 69
	fmt.Println(newMap["Deez Nutz"])

	newMap2 := map[string]string{
		"Ligma":     "Balls",
		"Sugandese": "Nutz",
	}

	fmt.Println(newMap2["Ligma"])
	if _, ok := newMap2["Sugandesez"]; ok {
		fmt.Println("Sugandese Nutz")
	} else {
		fmt.Println("Not Found")
	}

	for i, v := range newMap2 {
		fmt.Println(i, ":", v)
	}

	delete(newMap2, "Ligma")
	fmt.Println(newMap2)
}
