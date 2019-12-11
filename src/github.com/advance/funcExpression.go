package main

import "fmt"

func main(){
	s1 := foo()
	fmt.Println(s1)

	x := bar()

	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)


	fmt.Println(bar()())
}



func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}