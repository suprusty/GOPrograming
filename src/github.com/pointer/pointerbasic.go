package main

import "fmt"

func main() {
	x := 100
	fmt.Println("Valsu of x is ", x)
	fmt.Println("Address of x is ", &x)

	fmt.Printf("%T \n ", x)
	fmt.Printf("%T \n ", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(*&x)
	fmt.Println(&*y)

	*y = 1234007
	fmt.Println(x)


	val := 0
	fmt.Println("val befor", &val)
	fmt.Println("val befor", val)
	foo2(&x)
	fmt.Println("val after", &val)
	fmt.Println("val after", val)
}
func foo2(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}