package main

import "fmt"

func main() {
	fmt.Println("Welcome to IF-ELSE in Golang")
	/*
	 if...else Statement The if....else statement allows you to execute one block of code
	 if the specified condition is evaluates to true and another block of code if it is evaluates to false.
	*/
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "REGULAR USER"
	} else if loginCount > 10 {
		result = "HACKER"
	} else {
		result = "SOMETHING ELSE"
	}

	fmt.Println(result)

	//WE CAN INITIALIZE IN IF STATEMENT AS WELL in GOLANG
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}
