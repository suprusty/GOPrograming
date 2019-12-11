package main

import "fmt"

func main(){
	foo()
	defer bar()
	mat()
}

func foo(){
	fmt.Println("Foo ")
}

func bar(){
	fmt.Println("bar ")
}

func mat(){
	fmt.Println("mat ")
}
