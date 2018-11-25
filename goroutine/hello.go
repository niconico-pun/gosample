package main

import (
	"fmt"
	"sync"
)

func main() {
	go sayhello()

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "gretings", "good night"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func sayhello() {
	fmt.Println("hello")
}
