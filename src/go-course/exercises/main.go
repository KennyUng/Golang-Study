package main

import "fmt"

func main() {
	// Exercise 1
	/*
		arr := [5]int{}
		for i := 0; i < len(arr); i++ {
			arr[i] = i
		}

		for i, v := range arr {
			fmt.Println("Index", i, ":", v)
		}
	*/

	// Exercise 2
	/*

		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i, v := range slice {
			fmt.Println("Index", i, ":", v)
		}
	*/

	// Exercise 3
	/*
		slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		newSlice1 := slice[:5]
		newSlice2 := slice[5:]
		newSlice3 := slice[2:7]
		newSlice4 := slice[1:6]

		fmt.Println(newSlice1)
		fmt.Println(newSlice2)
		fmt.Println(newSlice3)
		fmt.Println(newSlice4)
	*/

	// Exercise 4
	/*
		slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		slice = append(slice, 52)

		fmt.Println(slice)

		slice = append(slice, 53, 54, 55)
		fmt.Println(slice)

		y := []int{56, 57, 58, 59, 60}
		slice = append(slice, y...)
		fmt.Println(slice)
	*/

	// Exercise 5
	/*
		slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		slice = append(slice[:3], slice[6:]...)
		fmt.Println(slice)
	*/

	// Exercise 7
	/*
		slice := [][]string{
			{"James", "Bond", "Shaken, not stirred"},
			{"Miss", "Moneypenny", "Hello James"},
		}
		for _, row := range slice {
			for i, v := range row {
				fmt.Println(i, v)
			}
		}
	*/

	// Exercise 8

	newMap := map[string][]string{
		"Bond":       {"James", "Martini", "Shaken, not stirred"},
		"Moneypenny": {"Miss", "James Bond", "Literature"},
		"No":         {"Doctor", "Being Evil", "Ice Cream"},
	}

	for i, v := range newMap {
		fmt.Println(i, ":", v)
	}
}
