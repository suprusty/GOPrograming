package main

import "fmt"

type person struct {
	first string
	last string
	age int
	ltk bool
}

type secretAgent struct {
	person
	ltk bool
}

func main(){
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
			ltk: false,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk,sa1.person.ltk)
	fmt.Println(p2.first, p2.last, p2.age,p2.ltk)
}