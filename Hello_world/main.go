package main

import (
	"fmt"
)

func main() {
	x := 100

	if x == 100 {
		fmt.Println(x, "x equals 100")
	} else if x > 100 {
		fmt.Println(x, "x greater than 100")
	} else {
		fmt.Println(x, "x less than 100")
	}

	//or => ||
	//and => &&
}
