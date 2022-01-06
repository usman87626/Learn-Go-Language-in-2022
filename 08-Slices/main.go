package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	   =======================SLICES========================
	   	The most used & powerful Datatype in Golang.
	   	Extended from Arrays.
	*/
	fmt.Println("Welcome to the Slices in Golang")

	/*
		================SYNTAX 01=========================
		In this syntax, it is necessary to initialize the Slice.
		In slices, we don't define how many values we are going to store.
	*/
	var fruitList = []string{"Apple", "tomato", "Peach"}
	fmt.Printf("Type of Fruitlist is %T \n", fruitList)

	/*
		IN CASE OF ARRAYS, WE ADDED VALUE USING INDEX: fruitList[0] = "VALUE"
		We can also do that here but Slices have much more optimized way to insert & delete values.
		i.e. append(slice, values)
		Let's use this in example => sliceVar = append(sliceVar, value1,value2,.....,valuen)
	*/
	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("Fruit List has: ", fruitList)

	//Slice Slicing Just like string/list slicing in Python
	fruitList = append(fruitList[1:4])
	fmt.Println("Sliced Slice is: ", fruitList)

	/*
		=================ANOTHER WAY TO CREATE A SLICE=====================
		In the Go programming language, the make() is a built-in function that is used to allocate
		and initializes an object of type slice, map, or chan (only).
		The return type of the make() function is the same as the type of its argument, not a pointer to it.

		make([]datatype,size)
	*/
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 NOT ALLOWED(ERROR - AS 4 is the Total Size)

	//append() Re-Allocates the memory - Dynamic Addition of Values
	highScores = append(highScores, 555, 666, 342)
	fmt.Println("High Scores: ", highScores)

	//SORT METHOD FOR SLICES
	sort.Ints(highScores)

	fmt.Println("Sorted highScore Slice: ", highScores)

	//CHECK IF SORTED
	fmt.Println("Is the highScore slice sorted: ", sort.IntsAreSorted(highScores))

	// DELETING VALUES FROM a SLICE based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "cpp", "ruby", "golang"}
	fmt.Println("Courses: ", courses)

	//INDEX THAT WE WANT TO REMOVE
	var index int = 2

	// courses[:2] => reactjs,javascript
	// courses[index+1:] => python cpp ruby golang
	// We will see ... in upcoming tutorials
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("UPDATED COURSES: ", courses)
}
