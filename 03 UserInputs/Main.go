package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to GoLang"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	//comma ok|| error ok d syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input)
	fmt.Printf("Hello %T", input)

	//storing error in a variable
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}
}
