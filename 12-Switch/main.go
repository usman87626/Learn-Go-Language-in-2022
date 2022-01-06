package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch-Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is: ", diceNumber)

	/*
		=========================SWITCH=================================
		It's similar to if-else.
		=> if-else better for boolean values: If-else conditional branches are great for
			variable conditions that result into a boolean, whereas
		=> switch statements are great for fixed data values. Speed: A switch statement might
			prove to be faster than ifs provided number of cases are good.

		In Golang, there is no need to specify `break` after each case, golang will do it for you
		but you can use `fallthrough` keyword in the end of Case if you want to run \
		all Cases below a Case that holds true.
	*/
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 & You can Open")
	case 2:
		fmt.Println("You can move 2 spots")
		fallthrough //When somebody hits Case 2, every case below this will also run
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots")
	default: //EXECUTED WHEN NO OTHER CASE HOLDS TRUE
		fmt.Println("What was that?")
	}
}
