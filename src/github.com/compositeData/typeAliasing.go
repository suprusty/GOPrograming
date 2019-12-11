package main

import "fmt"

type X struct{}

func (X) f() {}

type Y X // type declaration. cannot call functions on X via Y

type Z = X // alias declaration. identical types.

func main() {
	x := X{}
	y := Y{}
	z := Z{}

	fmt.Println(x, y, z)
	x.f()
	z.f()

	// y.f() // not possible
}