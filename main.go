package main

import (
	"fmt"
)

func main() {
	var dama int = 30
	println(dama, "\n")

	if dama <= 10 {
		fmt.Println("Weather is so cold")
	} else if dama > 10 && dama < 30 {
		fmt.Println("Weather is so good")
	} else {
		fmt.Println("Weather is so hard")
	}
}
