package main

import "fmt"

func main() {
	var names = []string{"Reza", "Hossein", "Mohamad", "Ali", "Hadi"}

	for index, name := range names {
		if name == "Ali" {
			names[index] = "Ali2"
			break
		}
	}

	fmt.Println(names)
}
