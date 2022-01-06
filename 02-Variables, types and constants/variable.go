package main

import "fmt"

/*
================CONSTANT========================
Once Initialized, we cannot change value of constant later in the program.
First charactetr of variable is Capital which is equivalent to public(can be accessible to any file or package in folder)
*/
const LoginToken string = "sd2d32d34nm4hj3h4343"

func main() {
	/*
		===============VARIABLE SYNTAX 1==========
		=> 	var variable_name datatype = value
	*/
	var username string = "Golang Variable"
	fmt.Println(username)
	//%T in Printf() is used to know the type of variable, similarly %v to print value of a variable.
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	// https://go.dev/ref/spec#Numeric_types
	var smallVal uint8 = 243 //Range between 0-255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type: %T \n", smallVal)

	// https://go.dev/ref/spec#Numeric_types
	var smallFloatVal float64 = 23.3242342342343 //the set of all IEEE-754 64-bit floating-point numbers
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of Type: %T \n", smallFloatVal)

	/*
		============DEFAULT VALUES===================
		Default value of a variable is a value that is present in a variable when we don't assign any value to it
		or when there is no value, there is a default value.

		=> Default Value for integer = 0 in golang
		=> Default Value for string  = empty white space
		=> Default Value for pointer = <nil> (We will see in upcoming code snippet)
	*/
	// Default Values & Some aliases
	var anotherVariable int //declaration - Default Value = 0 for int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherVariable)

	var anotherStringVariable string //declaration - Default Value = empty string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherStringVariable)

	/*
		======================ALIASES===========================
		Lexer of golang is so intelligent that If you initialize a variable on the spot when you were creating the variable,
		Lexer will automatically detect the type of the data & hence no explicity definition of datatype is needed in this case.
			Syntax => var variableName = value
		Depending on the value, Lexer will set variable to that datatype.

		NOTE: "We cannot just declare without initialization without declaring datatype"
		i.e. var variableName is not allowed, either initialize it on the spot otherwise use the old syntax as `var varName type`
	*/
	//Aliases
	// If we don't declare type while initialzing, lexer will infer type from our value
	// But we cannot just declare without declaring type
	var website = "google.com"
	fmt.Println(website)

	/*
		====================No var/type Style=====================
		In golang, the most widely used notation to create variables is using `:=` operator.
			Syntax => variableName := value
		NOTE: "You cannot use `:=`  operator for global variables"
	*/
	//no var style
	// We cannot use := operator outside of the method(global etc)
	numberOfUser := 33000
	fmt.Println(numberOfUser)

	numberOfUser = 34234
	fmt.Println(numberOfUser)

	fmt.Println("Login Token: ", LoginToken)
	fmt.Printf("Type of LoginToken: %T \n", LoginToken)

}
