package main

import "fmt"

func main() {
	fmt.Println("Welcome to the if else study")
	loginCount := 10

	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "New User"
	} else {
		result = "Unknown User"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	}

}
