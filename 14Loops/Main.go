package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops study")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days of the week: ", days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and the value is %v\n", index, day)
	}

	rouguevalue := 1
	for rouguevalue < 10 {

		if rouguevalue == 2 {
			goto gototest
		}
		if rouguevalue == 5 {
			rouguevalue++
			continue
		}
		fmt.Println("value is", rouguevalue)

		rouguevalue++

	}

gototest:
	fmt.Println("I am learning Go")
}
