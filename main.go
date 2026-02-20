package main

func main() {
	println(romanToInt("MCMXCIV"))

}

func romanToInt(s string) int {
	var sum byte = 0

	var letters = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if len(s) >= i+2 && letters[rune(s[i])] < letters[rune(s[i+1])] {
			sum -= byte(letters[rune(s[i])])
		} else {
			sum += byte(letters[rune(s[i])])
		}
	}

	return int(sum)
}
