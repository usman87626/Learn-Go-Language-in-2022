package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")
	//Let's say, we have some text/content
	content := "A QUICK BROWN FOX JUMPS OVER THE LAZY DOG A QUICK BROWN FOX JUMPS OVER THE LAZY DOG A QUICK BROWN FOX JUMPS OVER THE LAZY DOG A QUICK BROWN FOX JUMPS OVER THE LAZY DOG A QUICK BROWN FOX JUMPS OVER THE LAZY DOG "

	//os.Create(filename) will be used to create a file; Here we will create a .txt file
	//This operation can create file or throw an error, so we will use comma ok syntax
	file, err := os.Create("./myfile.txt")
	//Check if there is any error
	if err != nil {
		//panic shutdowns the program execution & show us the error
		panic(err)
	}

	//If the file is created successfully, io.WriteString(filename,content/data) is used to write String data inside the given file
	// Comma ok syntax is used here as we can get errors as well
	length, err := io.WriteString(file, content)

	//SAME AS ABOVE
	if err != nil {
		panic(err)
	}

	//Let's check the length of the written data(No. of characters)
	fmt.Println("Length is: ", length)
	// We will use defer here, the reason is that we will write more code below
	// Instead of moving file.Close() in the end or saying my friend to write everything upon file.Close(), it's good to use defer
	defer file.Close()

	//Calling the function that will read content of the file
	readFile("./myfile.txt")

}

//Function to Read Content from file
func readFile(filename string) {
	//To read file, we will use ioutil
	// ioutil will return dataByte or error, so we will use comma ok syntax
	dataByte, err := ioutil.ReadFile(filename)

	//SAME AS ABOVE
	if err != nil {
		panic(err)
	}
	//to convert buyte into regular data pass it into string(), we will see another methods to do as well in upcoming code snippets.
	fmt.Println("Text Data Inside the file is: ", string(dataByte))
}
