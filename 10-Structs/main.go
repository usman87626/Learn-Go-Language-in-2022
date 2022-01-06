package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Basics of Struct in Golang")
	/*
	   =============STRUCTS=====================
	   Like other languages have classes, Golang has no classes, Struct is the alternative here and are super useful.
	   Similarly, there is no Inheritance or super/parent concept in Golang.
	   Golang believes that when there is inheritance or super/parent stuff, code becomes complex to read
	*/
	usman := User{"USMAN", "usman.arif20188@gmail.com", true, 25}
	fmt.Println("STRUCT usman: ", usman)
	//FORMATTING: +v will give more details about struct (METADATA & DATA)
	fmt.Printf("usman details are: %+v\n", usman)
	//WE CAN GET SINGLE VALUE AS WELL
	fmt.Printf("Name is: %v, Email is: %v\n", usman.Name, usman.Email)
}

//DEFINING A STRUCTURE
// Capital U denotes its public & can be used anywhere,
// Similarly Name,Email & Status are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
