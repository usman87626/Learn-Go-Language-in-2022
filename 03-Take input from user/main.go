package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	/*
		Initializing the reader from a builtin package called bufio, os module will provide the address of the standard
		input device i.e. Keyboard
		os package => Used to specify the Standard Input Device (Keyboard) as os.Stdin
	*/
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the rating for our Pizza: ")

	/*
		===================COMMA OK SYNTAX / ERROR OK SYNTAX===================
		In Golang, there is no try catch exception handling, so this operation is perfromed using COMMA OK syntax
		we can receive error using comma separated value => input,error = someBuiltinModule
		If program executes properly, result will be assigned to input variable
		Otherwise, the error will be stored in `err` variable.

		NOTE: `If you don't care about err variable & want to optimize code, you can use the symbol `_`` instead of err`
	*/
	//ReadString until \n appears or In simple words, Until User press the Enter button
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	/*
		Type of the value in `input`` is string, we have to convert it into integer to perform numeric calculations.
		We will see Type-Casting/Type Conversion in next code snippet
	*/
	fmt.Printf("Type of this rating is %T \n", input)

}
