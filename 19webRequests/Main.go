package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to the Web Requests study")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %T \n", response)
	fmt.Println("Response is ", response)

	if response.StatusCode == 200 {
		fmt.Println("Response is OK")
	} else {
		fmt.Println("Response is not OK")
	}
	defer response.Body.Close()
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is ", string(dataBytes))
}
