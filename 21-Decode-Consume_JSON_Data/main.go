package main

import (
	"encoding/json"
	"fmt"
)

/*
	In the last part, we encoded/created JSON data from a Struct.
	It's time to decode/consume JSON data which is reverse process of encoding.
	We will copy the Encoding code from previous code snippet & create new function to decode it.
*/

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Consume/Decode JSON Data")
	//Call any function you want
	DecodeJsonAsMap()
	//DecodeJsonAsStruct()
}

/*
We will deocde it using 2 use-case.
	1) Deocde to Struct
	2)Decode to Map(key-value pairs)
*/

/*
1.Function that will Decode JSON Data to Struct format
*/
func DecodeJsonAsStruct() {

	//As data from Web is received as byte, so we will store it in byte
	jsonDataFromWeb := []byte(`
		{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
		}
	`)

	//Struct that will store the decoded data
	var myCourse course

	//Check if the JSON data is valid
	checkValid := json.Valid(jsonDataFromWeb) //returns true or false

	//IF Data is in Valid JSON format
	if checkValid {
		fmt.Println("JSON was valid")
		//json.Unmarshal() is used to decode JOSN. It take 2 arguments. 1)Data to Decode 2)Target variable to store decoded value
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		// `%#v` will give more detail of the output
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("There is something wrong with JSON data")
	}
}

/*
2.Function that will decode JSON data to Map
*/
func DecodeJsonAsMap() {
	//As data from Web is received as byte, so we will store it in byte
	jsonDataFromWeb := []byte(`
		{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
		}
	`)

	//we can use datatype as `interface` when we don't know exactly what type we are going to store in it
	var myMapData map[string]interface{}
	//json.Unmarshal() is used to decode JOSN. It take 2 arguments. 1)Data to Decode 2)Target variable to store decoded value
	json.Unmarshal(jsonDataFromWeb, &myMapData)

	for k, v := range myMapData {
		fmt.Printf("Key is %v, Value is %v and Type is: %T\n", k, v, v)
	}
}
