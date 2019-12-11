package main

import (
	"fmt"
)

type T1 struct{}
type T2 struct{}

type T3 struct {
	i int
	s string
}

type T4 struct {
	i  int
	sl []int
}

type T5 struct {
	i  int
	ar [3]int
}

func main() {
	t1 := T1{}
	t2 := T2{}
	//fmt.Println("t1 == t2", t1==t2) // invalid operation: t1 == t2 (mismatched types T1 and T2)
	_, _ = t1, t2

	t1a, t1b := T1{}, T1{}
	fmt.Println("t1a == t1b", t1a == t1b) // true

	t3a, t3b, t3c := T3{5, "hello"}, T3{5, "hello"}, T3{10, "hello"}
	fmt.Println("t3a == t3b", t3a == t3b) // true
	fmt.Println("t3b == t3c", t3b == t3c) // false

	t4a, t4b := T4{1, []int{}}, T4{1, []int{}}
	//fmt.Println("t4a == t4b", t4a == t4b) // invalid operation: t4a == t4b (struct containing []int cannot be compared)
	_, _ = t4a, t4b

	t5a, t5b, t5c := T5{1, [3]int{1, 2,3}}, T5{1, [3]int{1, 2,3}}, T5{1, [3]int{11, 12,13}}
	fmt.Println("t5a == t5b", t5a == t5b) // true
	fmt.Println("t5b == t5c", t5b == t5c) // false
}