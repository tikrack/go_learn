package main

func main() {
	text := "   Hello my name is tikrack   "
	runes := []rune(text)

	a := false
	length := 0

	for i := len(runes) - 1; i >= 0; i-- {
		char := runes[i]

		if !a && char == ' ' {
			continue
		}
		a = true
		if char == ' ' {
			break
		}
		length++
	}

	println(length)
}
