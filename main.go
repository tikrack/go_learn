package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList []string

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go Get("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(i), &wg)
	}

	wg.Wait()

	fmt.Println(TodoList)
}

func Get(url string, wg *sync.WaitGroup) {
	println(url)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	TodoList = append(TodoList, string(res))

	if err != nil {
		panic(err)
	}

	defer wg.Done()
}
