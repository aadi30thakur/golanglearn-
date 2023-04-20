package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")
	defer mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("hello", i)
	}

}
