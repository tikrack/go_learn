package main

import "fmt"

func c() {
	fmt.Print("\033[H\033[2J")
}

const (
	Reset  = "\033[0m"
	Yellow = "\033[33m"
)

func main() {
	var project byte = 1
	var run = false

	for run == false {
		c()

		println(Yellow)
		println(" ╔═════╦════════════════════════════════╗")
		println(" ║ ID  ║ Project                        ║")
		println(" ╠═════╬════════════════════════════════╣")

		if project == 1 {
			print(" >")
		} else {
			print(" ║")
		}
		fmt.Printf("  1  ║ %s                 ║\n", "Throw the dice")

		println(" ╚═════╩════════════════════════════════╝")
		println(Reset)

		run = true
	}
}

func one() {
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
