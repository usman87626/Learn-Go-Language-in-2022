package main

import "fmt"

func main() {
	/*
	   	===================POINTERS==================
	   There is a problem in every problem language i.e. Sometimes we declare variable or constant
	   They get stored in some memory, when we pass them to some function or anywhere, their copies
	   are passed and we face different problems and irregularities in our apps.
	   To overcome it, we use pointers that have the exact address of the location.
	   Pointers provide us the surety that whatever we are passing on, it is always the
	   actual value from the memory.
	*/
	fmt.Println("Welcome to a Class on Pointers")

	// DEFAULT VALUE is <nil>
	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)
	/*
		Default Value in pointer is <nil>
		Accessing pointer using * as *pointerName is equivalent to using variableName that the pointer is referring to.
		In the below example, second method to write myNum is *ptr as ptr has address of myNum(&myNum)
	*/
	myNum := 23
	var ptr = &myNum

	fmt.Println("Value of actual pointer is: ", ptr)
	fmt.Println("Value inside the pointer is: ", *ptr)
	fmt.Println("Address of actual variable: ", &myNum)

	*ptr = *ptr * 2
	fmt.Println("New Value inside the pointer is: ", *ptr)
	fmt.Println("New Value inside the variable is: ", myNum)

}
