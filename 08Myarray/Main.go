package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array study")
	var arry [4]string
	arry[0] = "Hello"
	arry[1] = "World"
	arry[2] = "!"
	arry[3] = "!"
	fmt.Println(arry)
	fmt.Println(len(arry))

	var VegList = [3]string{"Potato", "Tomato", "Onion"}

	fmt.Println("veggieList", VegList, len(VegList))
}
