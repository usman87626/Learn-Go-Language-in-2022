package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
We will utilize http package to understand basics of web requests
More info here: https://pkg.go.dev/net/http
*/

//WEBSITE WE ARE GOING TO SEND REQUEST AND GET ITS HTML
const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web Requests in Golang")

	response, err := http.Get(url)

	//Error Checking
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of Response: %T \n", response) //*http.Response type

	//It's our responsibility to close the connection everytime, as suggested by documentation
	defer response.Body.Close()

	//Reading Response using ioutil package
	databyte, err := ioutil.ReadAll(response.Body)

	//Error checking
	if err != nil {
		panic(err)
	}
	//Converting databyte to string
	content := string(databyte)

	fmt.Println(content)
}
