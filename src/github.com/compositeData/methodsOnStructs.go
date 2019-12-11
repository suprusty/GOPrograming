package main

import "fmt"

type rect struct{
	width,height int
}


func (r rect) perim() int{
	return 2* r.height + 2*r.width
}

func (rpointer *rect) area() int{
	return rpointer.width * rpointer.height
}

func main(){
	rs := rect{10,20}
	fmt.Println(rs.perim())
	fmt.Println(rs.area())

	rsPointer := &rs
	fmt.Println(rsPointer.perim())
	fmt.Println(rsPointer.area())
}