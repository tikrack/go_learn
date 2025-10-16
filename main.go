package main

func main() {
	var dama = 20
	println(dama, "\n")

	switch dama {
	case 0:
		println("Hava is zero 0")
	case 1:
		println("Hava is not zero 1")
	default:
		println("Hava is not zero 2")
	}

	switch {
	case dama <= 0:
		println("Hava is zero 0")
	case dama > 10 && dama < 100:
		println("Hava is not zero 1")
	default:
		println("Hava is not zero 2")
	}

}
