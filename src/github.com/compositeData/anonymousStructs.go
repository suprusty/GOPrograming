package main

import "fmt"

// note that this does not use a type keyword
var config struct {
	APIKey     string
	ConnString string
	db         int //db connection handle
}

func main() {
	config.ConnString = "user:password@tcp(hostname:3036)"
	fmt.Println("conn string is: ", config.ConnString)

	s := struct {
		name string
		age int
	}{
		"john",
		32,
	}
	fmt.Println(s)
}
