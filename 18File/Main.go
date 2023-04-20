package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to the Files study")
	content := "This is the content of the file"
	file, err := os.Create("test.txt")

	checkNilError(err)
	if err != nil {

		panic(err)
	}
	length, err := io.WriteString(file, content)
	checkNilError(err)
	// if err != nil {

	// 	panic(err)
	// }
	fmt.Println("File created with length ", length)
	readFile("./test.txt")
	defer file.Close()
}

func readFile(fileName string) {

	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("File content is \n", string(dataByte), dataByte)
}

func checkNilError(err error) {

	if err != nil {
		panic(err)
	}
}
