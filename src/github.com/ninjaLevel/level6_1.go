package main

import "fmt"

func main(){
 f := foo1()
 fmt.Println(f())


	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := callback(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)


	ec := encloses()
	fmt.Println(ec())
	fmt.Println(ec())
	gc := encloses()
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
}



func foo1() func() int{
	return func() int{
		return 200
	}
}


func callback(f func(xi []int) int,ii []int) int{
	n := f(ii)
	n++
	return n
}

func encloses()  func() int{
	globaScope := 100
	return func() int{
		globaScope++
		return globaScope
	}
}