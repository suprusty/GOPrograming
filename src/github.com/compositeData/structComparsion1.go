package main

import (
	"fmt"
)

type personI interface {
	age() int
}
type dev struct {
	exp int
}

func (p dev) age() int {
	return p.exp+20
}

type ops struct {
	exp int
}

func (o ops) age() int {
	return 6
}

type computer struct {
}

func main() {
	var d1 = dev{10}
	var d2 = dev{10}
	fmt.Println(d1 == d2) // true

	var p1 personI = d1
	var p2 personI = d2
	fmt.Println(p1 == p2) // true

	var p3 personI = ops{10}
	fmt.Println(p1 == p3) // false: not same underlying type but ...
	// they can be compared as they implement ...
	// the same interface; not a compiler error

	var c = computer{}
	//fmt.Println(p1 == c) // compiler error: invalid operation:  ...
	// p1 == c (mismatched types person and computer)
	_ = c
}