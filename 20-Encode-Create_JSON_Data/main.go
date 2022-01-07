package main

import (
	"encoding/json"
	"fmt"
)

/*
=============GOAL OF THIS CODE SNIPPET================
Encoding data to JSON format

======================LETS CREATE A STRUCT==============
Lets create a Struct of Type course(not public)

=> We can add `json:"coursename"` with Name string & when data is converted to JSON, it will renamed `Name` to "coursename"
        {
                "coursename": "PHP Bootcamp",
                "Price": 199,
                "website": "LearnCodeOnline.in"
        }
		Although, real name of the attribute in Struct Declaration was 'Name' but we overwritten it by "coursename"

=> If there is some field, that you don't want to show after encoding data to JSON, you can use `json:"-"` for that,
   as we did for 'Password', run the code after removing it once & you will know why it's here.

=> Similarly,if there is a field that possibly can be Null sometimes & we don't want to show it as "tags": nil
   Then we can use "omitempty" in the above syntax as we did in the case of 'Tags' -> `json:"tags,omitempty"`

=> Remove these literals & run the code. Then put back these literals & run the code & you will understand it more precisely.
*/
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON CREATION in Golang")
	//Let's call the function
	EncodeJson()
}

func EncodeJson() {
	//Let's Create Slice of our Struct
	myCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "mypass123", []string{"web-dev", "js"}},
		{"Python Bootcamp", 255, "LearnCodeOnline.in", "usmanarif", []string{"python-dev", "python"}},
		{"Golang Bootcamp", 388, "LearnCodeOnline.in", "123golangdev", []string{"golang-dev", "go"}},
		{"PHP Bootcamp", 199, "LearnCodeOnline.in", "phpdev3321", nil},
	}

	/*
		==========Lets encode it to JSON===================
		json.Marshal() will be used to do it. It accepts Struct(interface).
		WE can also use json.MarshalIndent() which will give the indentated data.
		MarshalIndent() gets 3 parameters: 1)data 2)String that should be added to each line 3)\t to give space of tab

	*/
	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}
