package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch study")
	// loginCount := 10

	// switch loginCount {
	// case 10:
	// 	fmt.Println("Regular User")
	// case 11:
	// 	fmt.Println("New User")
	// default:
	// 	fmt.Println("Unknown User")
	// }

	// switch 9 % 2 {
	// case 0:
	// 	fmt.Println("9 is even")
	// default:
	// 	fmt.Println("9 is odd")
	// }

	// switch num := 3; num < 10 {
	// case true:
	// 	fmt.Println("Number is less than 10")
	// }

	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	fmt.Println("Dice value is", dice)
	switch dice {
	case 1:
		fmt.Println("You got 1")
	case 2:
		fmt.Println("You got 2")
		fallthrough
	case 3:
		fmt.Println("You got 3")
		fallthrough
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	default:
		fmt.Println("You got nothing")
	}
}
