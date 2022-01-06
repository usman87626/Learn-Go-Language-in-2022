package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	//Let's Create A Slice
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	//Printing All the Values
	fmt.Println("Values in days Slice: ", days)

	/*
		==========SYNTAX 1=============================
		SYNTAX: for counter_init; condition_to_break; increment {body}
		There is no pre-increment in Golang ++counter or ++d etc

		This syntax resembles the C/C++ syntax of loop => for(counter_init; condition; increment){body}
	*/
	fmt.Println("====SYNTAX 1====")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	/*
		==========SYNTAX 2=============================
		SYNTAX: for placeholder := range slice/variable {BODY}
	*/
	fmt.Println("====SYNTAX 2====")
	for i := range days {
		fmt.Println(days[i])
	}
	/*
		==========SYNTAX 3=============================
		It's something like For Each
		SYNTAX: for index,value := range slice/variable {BODY}
	*/
	fmt.Println("====SYNTAX 3====")
	for index, value := range days {
		fmt.Printf("INDEX is %v, VALUE is %v \n", index, value)
	}

	/*
		BREAK & CONTINUE STATEMENT:

		BREAK => BREAK THE LOOP/BODY
		CONTINUE => SKIPS THE ITERATION
	*/
	fmt.Println("BREAK & CONTINUE STATEMENT")
	randomValue := 1
	for randomValue < 10 {
		if randomValue == 5 {
			break
		}

		if randomValue == 2 {
			goto mytag
		}
		fmt.Println("VALUE is: ", randomValue)
		randomValue++
	}

	/*
		goto in GOLANG
	*/
mytag:
	fmt.Println("INSIDE GOTO TAG")

}
