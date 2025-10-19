package main

import "fmt"

func main() {
	var adads = make([]int, 4, 6)

	fmt.Println("Len: ", len(adads))
	fmt.Println("Cap: ", cap(adads))
	fmt.Println("----------------------------------------")

	adads = append(adads, 0, 2)

	fmt.Println("Len: ", len(adads))
	fmt.Println("Cap: ", cap(adads))
	fmt.Println("----------------------------------------")

	adads = append(adads, 0, 2)

	fmt.Println("Len: ", len(adads))
	fmt.Println("Cap: ", cap(adads))
}
