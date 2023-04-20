package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices study")
	var fruitList = []string{}
	fmt.Printf("type of fruitList is %T \n", fruitList)
	fruitList = append(fruitList, "Apple", "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList1 := append(fruitList[1:])
	fmt.Println(fruitList1)

	fruitList1 = append(fruitList[1:2])
	fmt.Println(fruitList1)

	fruitList1 = append(fruitList[:2])
	fmt.Println(fruitList1)

	highScore := make([]int, 4)
	highScore[0] = 2134
	highScore[1] = 4123
	highScore[2] = 1234
	highScore[3] = 1241234
	// highScore[4] = 123555

	highScore = append(highScore, 123555, 123455)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courseLsit = []string{"Python", "Go", "Java", "C++", "C", "Ruby", "JavaScript"}
	fmt.Println(courseLsit)
	var index int = 2

	courseLsit = append(courseLsit[:index], courseLsit[index+1:]...)
	fmt.Println(courseLsit)

}
