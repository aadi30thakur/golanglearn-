package main

import "fmt"

const loginToken string = "1234567890" //public variable

func main() {
	var username string = "Aditya Thakur"
	fmt.Println(username)
	fmt.Printf("Type of username is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of username is %T \n", isLoggedIn)

	var small int = 256
	fmt.Println(small)
	fmt.Printf("Type of username is %T \n", small)

	var smallfloat float64 = 255.55675797676765795769796768976
	fmt.Println(smallfloat)
	fmt.Printf("Type of username is %T \n", smallfloat)

	var smallcomplex complex64 = 255.55675797676765795769796768976 + 255.55675797676765795769796768976i
	fmt.Println(smallcomplex)
	fmt.Printf("Type of username is %T \n", smallcomplex)

	//default value and alias

	var anothervariable int
	fmt.Println("Type of username is %T \n", anothervariable)

	var url = "https://www.google.com"
	fmt.Println(url)

	// no var style

	numberofusers := 100000
	fmt.Println(numberofusers)

	fmt.Println(loginToken)
	fmt.Printf("Type of username is %T \n", loginToken)
}
