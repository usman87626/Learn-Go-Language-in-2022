package main

import "fmt"

func main() {
	fmt.Println("Methods of Struct in Golang")
	/*
		===========METHODS=====================
		Methods is nothing but just a fancy name for functions.
		Functions when used with Struct are termed as methods.
	*/

	usman := User{"USMAN", "usman.arif20188@gmail.com", true, 25}
	fmt.Println("STRUCT usman: ", usman)
	//FORMATTING: +v will give more details about struct (METADATA & DATA)
	fmt.Printf("usman details are: %+v\n", usman)
	//WE CAN GET SINGLE VALUE AS WELL
	fmt.Printf("Name is: %v, Email is: %v\n", usman.Name, usman.Email)
	usman.GetStatus()
	usman.NewMail()
	fmt.Printf("Name is: %v, Email is: %v\n", usman.Name, usman.Email)
}

//DEFINING A STRUCTURE
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

/*
=============DECLARING A METHOD====================
Syntax for declaring a method is:
func (object Struct) function_Name(arg1,arg2,arg3,...) return_type {

}

=> In this code snippet, we will create GetStatus() which will print status of the user
=> Similary, we will use NewMail() to manipulate email of user but it will pass the copy to make changes,
	which means that value will not be changed permanently, to change it permanently we will pass struct as pointer whihc we will discuss
	in upcoming code snippets.
*/
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newmail@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
