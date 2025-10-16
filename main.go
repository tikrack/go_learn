package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello my name is tester"
	var splited = strings.Split(str, " ")
	fmt.Println(strings.Join(splited, ","))
}
