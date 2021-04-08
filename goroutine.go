package main

import "fmt"

func main() {
	hello()
	fmt.Println("main goroutine done!")
}

func hello() {
	fmt.Println("Hello Goroutine")
}
