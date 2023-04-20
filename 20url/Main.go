package main

import (
	"fmt"
	"net/url"
)

//create a rsndom url string

const myurl string = "https://lco.dev:3000/learn?coursename=ReactJS&paymentid=free"

func main() {
	fmt.Println("Welcome to the Path study")
	neturl, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(neturl.Scheme)
	fmt.Println(neturl.Host)
	fmt.Println(neturl.Path)
	fmt.Println(neturl.RawQuery)
	fmt.Println(neturl.Port())
	fmt.Println(neturl.RequestURI())

	qparams := neturl.Query()
	fmt.Printf("the type of qparams %T\ns", qparams)
	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "coursename=ReactJS&paymentid=free",
	}

	anotherUrl := partsofURL.String()
	fmt.Println(anotherUrl)
}
