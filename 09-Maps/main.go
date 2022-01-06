package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Maps in Golang")
	/*
	   ===================MAPS=======================
	   In some languages, its known as hash table or Key-Value Pair or dictionary.
	*/
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CPP"] = "C Plus Plus"

	fmt.Println("List of All Languages: ", languages)
	fmt.Println("JS stands for: ", languages["JS"])
	fmt.Println("PY stands for: ", languages["PY"])
	fmt.Println("RB stands for: ", languages["RB"])
	fmt.Println("CPP stands for: ", languages["CPP"])

	//Delete a Value from Map
	delete(languages, "RB")
	fmt.Println("Updated List of Languages: ", languages)

	//ADDITIONAL STUFF (LOOPS ARE INTERESTING in GOLANG)
	for key, value := range languages {
		fmt.Printf("For Key %v, Value is %v \n", key, value)
	}

}
