package main

import "fmt"

func main() {
	one()
}

func one() {
	const (
		Reset  = "\033[0m"
		Yellow = "\033[33m"
	)

	var n int

	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	fmt.Println(Yellow)

	for i := 1; i <= n; i++ {
		for x := n - i; x > 0; x-- {
			fmt.Print("  ")
		}

		for y := 1; y <= i; y++ {
			fmt.Printf(" %d", y)
		}

		for z := i - 1; z >= 1; z-- {
			fmt.Printf(" %d", z)
		}

		fmt.Println()
	}

	fmt.Println(Reset)
}
