package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs study")

	hitesh := User{"Aditya", "aditya@gmial.com", true, 23}

	fmt.Println(hitesh)
	fmt.Printf("User Details  are : %+v\n", hitesh)
	fmt.Printf("User Name is %v and email is : %v\n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()

	u := hitesh.NewMail()
	fmt.Println("New Email is ", u)
	fmt.Printf("User Name is %v and email is : %v\n", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// these are the methods one that can retun a value and other that can't
func (u User) GetStatus() {
	fmt.Println("Is user active ", u.Status)
}

func (u User) NewMail() string {
	u.Email = "test@go.dev"
	fmt.Println("New Email is ", u.Email)
	return u.Email
}
