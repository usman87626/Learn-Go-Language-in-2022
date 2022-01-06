package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	fmt.Println("Please rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type is: %T \n", input)
	/*
		During reader.ReadString('\n'), \n was added to the end of the string.
		First of all we have to remove \n to type cast string to integer as \n cannot be converted to a number.
		Removal of \n can be done as => strings.TrimSpace(input)

		After removal of \n, we can now use ParseFloat() method from strconv package as strconv.ParseFloat(strings.TrimSpace(input),64)

	*/
	// REMOVE \n first using strings.TrimSpace(input)
	//TYPE CONVERSION(TYPE CAST)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
