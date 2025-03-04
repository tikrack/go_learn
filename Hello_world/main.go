package main

import "fmt"

func main() {
	x := 100

	swich(x)
}

func swich(x int) {
	switch x {
	case 100:
		fmt.Println("x is 100")
		break
	case 200:
		fmt.Println("x equal to 200")
		break
	default:
		fmt.Println("swich fail")
	}
}
