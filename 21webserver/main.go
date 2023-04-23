package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformaGETRequest()
	preformPOSTRequest()
	preformPOSTFORMRequest()
}

func PerformaGETRequest() {
	const myURL string = "http://localhost:3000"
	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//print response body in string format

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(content))

	var responseString strings.Builder
	contentstr, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(contentstr)

	fmt.Println("Byte Count", byteCount)
	fmt.Println(response.StatusCode, responseString.String())
}

func preformPOSTRequest() {
	const myURL string = "http://localhost:3000/post"

	requestBody := strings.NewReader(`{
		"first_name": "John",
		"last_name": "Doe",
		"email": "aadi30.thakur@gmail.com"

	}`)
	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(content))
}

func preformPOSTFORMRequest() {
	const myURL string = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("first_name", "John")
	data.Add("last_name", "Doe")
	data.Add("email", "aadi30.thakur")
	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(content))

}
