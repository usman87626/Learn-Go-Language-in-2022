package main

import "fmt"

/*
	===================DEFER=====================
	 When a function executes, it executes line by line which is the basic flow
	 of each program either interpreted or compiled.When we put defer keyword,
	 whatever the next line we've marked for execution is not going to execute there
	 but will execute at the very end of the function.

	 => When we use defer in the last line of function, it has no effect.
	    The effect is maximum when written on the top portion of the function.
	 => In simple words, If we use defer Println("SOME STRING") anywhere in function,
	    defer keyword will cut that line and paste it before the ending braces of the function
	=> If there are so many deferred statements, Last one will be printed first of all and so on i.e. Last In First Out(LIFO)

*/
func main() {
	defer fmt.Println("FIRST")
	defer fmt.Println("SECOND")
	defer fmt.Println("THIRD")
	defer fmt.Println("FORTH")
	fmt.Println("Hello World")
	myDefer()

}

//Let's Visualize the Output
//First 4 lines will be placed in the last of program
// FIRST SECOND THIRD FORTH
// Hello World will be printed then
// Function will be called and loop is executed
// defer will delay the printing and 01234 will be queued
// Final output using LIFO is:
// Hello World 43210 FORTH\n THIRD\n SEOND\n FIRST\n

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
