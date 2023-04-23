package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to the JSON study")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 200, "lcolearn.code.online", "gasfdgasfdgadfg", []string{"react", "js"}},
		{"MernBootCamp", 199, "lcolearn.code.online", "fdsadasfads", []string{"FullStack", "fasdf"}},
		{"Angular", 299, "lcolearn.code.online", "afsdfasdfasdfsadgdfsa", nil},
	}
	finalJson, err1 := json.Marshal(lcoCourses) //need to send always the interface  ie struct
	if err1 != nil {
		panic(err1)
	}

	finalJson1, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
	fmt.Printf("%s\n", finalJson1)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
		             "coursename": "ReactJS",
		             "price": 200,
		             "platform": "lcolearn.code.online",
		             "tags": ["react","js"]
		     }`)

	var lcoCourse course

	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("Json was not valid")
	}

	//case where we just want to add data to key value
	var myOnlinedata map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlinedata)
	fmt.Printf("%#v\n", myOnlinedata)

	for key, value := range myOnlinedata {
		fmt.Printf("key is %v and value is %v and TYpe is %T\n", key, value, value)
	}

}
