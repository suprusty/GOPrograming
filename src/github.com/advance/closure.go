package main

import "fmt"

//one scope enclosing other scopes
var x int
func main(){
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	a:= incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("hello")
	x++
}

func boo(){
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y)

	fmt.Println(x)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}