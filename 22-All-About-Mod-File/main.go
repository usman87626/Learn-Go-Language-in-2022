package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
======GO MODULE(go mod command)===================

1. go mod init module-name
	module-name as myapp or calculator etc is not a good practice, it's better
	to use github.com/username or any hosting where we are going to put our code.
	This command will create go.mod file with two lines:
	module module-name
	//VERSION OF THE GO LANGUAGE
	go 1.17

We will use Routing System(Gorilla MUX) for this tutorial.
See MUX on github => https://github.com/gorilla/mux

2. Installing MUX package using the command:
	`go get -u github.com/gorilla/mux`
	It will pull all the dependencies, `go get` will go to web
	and bring the repository to our system(NOT IN THE PROJECT FOLDER)
	After this command, check the changes in `go.mod` file. It will add a new line:
			`require github.com/gorilla/mux v1.8.0 // indirect`
	The word indirect will be there until we are not using this package.
	Similary, after this command a new file is created `go.sum` which includes checksums for the installations.
	We can verify the installation using `go mod verify`` command & it will tell us if everything is good or not.

3. Run the command `go env` & copy value of GOPATH, this is the place where all the libraries & packages
	are downloaded. This is the reason, the MUX module was downloaded but not visible in our project folder.

4. Let's write some code & run it to see output on web page at address=> localhost:4000.
   Use go build . to build everything in same folder. Then run using go main.go

5. If you see go.mod file even after running the code, there is still `indirect` written to the right of module.
	Use `go mod tidy` to fix it.

==========SOME MORE COMMANDS=======================
1. We can use `go mod list` & it will list our package
2. Use `go mod all` to print every package/module
3. Use `go mod -m all` to print all the modules in same project/file.
4. Use `go list -m -versions github.com/gorilla/mux` will list all the available versions of 'MUX'
5. Use `go mod graph` & it will all the packages will all dependencies
6. Use `go mod vendor` to package everything in vendor folder(All dependencies,etc)
   Then we can run code using `go run -mod=vendor` main.go & it will run without loading packages
   from Internet or Cache of your 'GOPATH' where packages are installed.
7. Read about more commands here => https://go.dev/ref/mod
*/
func main() {
	fmt.Println("Mod File Basics in Golang")
	//FROM MUX documentation
	r := mux.NewRouter()
	//Binding route with our function & use `GET` method
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//FUNCTION TO WRITE SOME DATA on page WHEN MUX calls it
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Golang Developers!</h1>"))
}
