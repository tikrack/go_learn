package main

import "fmt"

func main() {
	//Array
	names := []string{"Ali", "Mohamad", "Reza", "Armin"}

	fmt.Println("Names: ", names)

	//Slice
	var onToTwo = names[1:3] //3=2+1

	fmt.Println("On To Two: ", onToTwo)
}
