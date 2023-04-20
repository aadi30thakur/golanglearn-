package main

import "fmt"

func main() {

	fmt.Println("Welcome to the maps study")
	languages := make(map[string]string)
	languages["Go"] = "Go is a programming language"
	languages["Python"] = "Python is a programming language"
	languages["Java"] = "Java is a programming language"
	languages["C"] = "C is a programming language"
	languages["C++"] = "C++ is a programming language"
	languages["C#"] = "C# is a programming language"
	languages["JavaScript"] = "JavaScript is a programming language"

	fmt.Println(languages)
	fmt.Printf("The length of the map is %d , value %v", len(languages), languages["java"])
	// Delete a key
	delete(languages, "C#")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("Key: %s , Value: %s, \n", key, value)
	}

}
