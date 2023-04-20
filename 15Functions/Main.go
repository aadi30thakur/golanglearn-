package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions study")
	greet := greater

	greet()

	result := adder(10, 20)
	fmt.Println("Result is  ", result)

	value, _ := proadder(10, 20, 30, 40, 50)
	val, p1 := proadder(10, 20, 30, 40, 50)
	fmt.Println("Value is ", value, val, p1)
}

func greater() {
	fmt.Println("Namaste from Golang")
}

func adder(a int, b int) int {
	return a + b
}

func proadder(values ...int) (int, string) {
	result := 0
	for _, v := range values {
		result += v
	}
	return result, "Hi pro res"
}
