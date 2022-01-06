package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Arrays in Golang")

	/*
		========================ARRAYS in GOLANG=========================
		In golang, arrays are the least used data-type due to their strict nature i.e. You have to specify exact number
		of elements in advance that you want to store in array & most of the time it's hard to tell how much elements we will
		store in a list/array. So, whenever you know the exact number of values, Array can be used there.
		Other datatype like Slice is also created using Array datatype.
	*/
	var fruitList [5]string
	/*
		==============INDICES=================================
		INDICES ARE USED TO ASSIGN/RETREIVE VALUES FROM DIFFERENT POSITIONS OF THE ARRAY.
		INDEX STARTS FROM 0 like other programming languages.
	*/
	fruitList[0] = "Apple"
	fruitList[1] = "Tomate"
	//fruitList[2] will get '' by default
	fruitList[3] = "Orange"
	fruitList[4] = "None"
	fmt.Println("Fruit List is: ", fruitList)
	//It will give us the total size of Array, Not the total number of elements in case of Arrays
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("Veg List is: ", vegList)
	fmt.Printf("Type of Veg List is %T \n", vegList)
}
