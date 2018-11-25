package main

import (
	"fmt"
)

func main() {
	go sayhello()
}

func sayhello() {
	fmt.Println("hello")
}
