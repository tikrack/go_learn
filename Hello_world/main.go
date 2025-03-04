package main

import "fmt"

func main() {
	var number int
	number = 100

	var float float64
	float = float64(number)

	fmt.Printf("value: %v and Type: %T \n", number, number)
	fmt.Printf("value: %v and Type: %T \n", float, float)
}
