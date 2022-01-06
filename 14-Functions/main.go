package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Functions in Golang")
	/*
		=======================FUNTIONS========================
		A function is a group of statements that exist within a program for the purpose of
		performing a specific task. At a high level, a function takes an input and returns an output.
		=> Function allows you to extract commonly used block of code into a single component.
		=> Functions cannot be created inside functions in Golang.
		=> Golang has anonymous functions as well, we will discuss it in upcoming code snippets.
	*/

	//Calling the function that we declared below
	greeter()

	result := adder(3, 4)

	fmt.Println("Sum is: ", result)

	proResult, ok := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Pro Result: ", proResult, ok)

}

/*
===========FUNCTION DECLARATION=======================
Syntax for function declaration:
func function_Name(argument1,argument2,argument3,...) return_type {
	BODY
}

=> We can use as many arguments as much we need or we can use none
=> We can return as many values as we want but we have to encapsulate it into parenthesis () becauase Functions as a whole can return
   only a single value. (In case of multiple values, all values are packed into a single component and treated as a single unit, see proAdder() for demo)

=> We can use ... to accept unlimited arguments in a function, this kind of function is called Variadic function.
*/

//FUNCTION WITH NO RETURN TYPE & NO ARGUMENTS
func greeter() {
	fmt.Println("Hello Greeter Function")
}

//Function with two arguments & int return type
func adder(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

//Variadic Function with unlimited arguments and (int,string) return type for total,"OK" message
func proAdder(values ...int) (int, string) {
	total := 0
	// Loop over slice to add values and save result in `total`
	for _, value := range values {
		total += value
	}
	// return the value of `total` back to the caller
	return total, "All OK"
}
