package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers study")

	//declare a variable
	// var myInt int
	// myInt = 42
	// fmt.Println("myInt value is", myInt)
	// fmt.Println("myInt address is", &myInt)

	var ptr *int
	print("ptr value is ", ptr)
	MyNumber := 42
	var ptr1 = &MyNumber
	fmt.Println("ptr1 value is ", ptr1)
	fmt.Println("ptr1 value is ", *ptr1)
	*ptr1 = *ptr1 * 100
	fmt.Println("ptr1 value is ", *ptr1)
}
