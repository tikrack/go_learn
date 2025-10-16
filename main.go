package main

func main() {
	var list = []string{"Hello", "How", "Are", "you"}

	var i = 0

	for i < 10 {
		i++
		println("Hello, World!")
	}

	for x := 10; x > 0; x-- {
		println(x)
	}

	for index, item := range list {
		if index == 2 {
			//break
			//continue
		}
		println(index, item)
	}
}
