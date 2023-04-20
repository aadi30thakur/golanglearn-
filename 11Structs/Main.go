package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs study")

	hitesh := User{"Aditya", "aditya@gmial.com", true, 23}

	fmt.Println(hitesh)
	fmt.Printf("User Details  are : %+v\n", hitesh)
	fmt.Printf("User Name is %v and email is : %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
