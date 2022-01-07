package main

import (
	"fmt"
	"net/url"
)

//Creating a Random URL to manipulate & get info from it
const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbdb2"

func main() {
	fmt.Println("Welcome to URL Manipulations in Golang")

	//Parsing URL with url library
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	//.Query() will give us Queries of URL in Map datatype format => url.Values
	qparams := result.Query()

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	/*
		============Let's do the reverse process===============
		Creating URL from different parts.
		We have to call &url.Url{}. Don't forget '&' sign
	*/
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=usman",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
