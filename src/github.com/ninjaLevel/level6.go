package main

import "fmt"

func main(){

	defer me()

	slX := []int{1,2,3,4,5,6,7,8,9}
	f := foo(slX...)
	fmt.Println(f)

	i,s := bar(slX)
	fmt.Println("String return ",s,"int return ",i)

	p := person{
		"sushil",
		"Kumar",
		40,
	}
	p.speak()
}

func foo(slx ...int) int {
	sum :=1000
	for _,x :=  range slx{
		sum +=x
	}
	return sum
}

func bar(slx []int) (int,string){
	sum := 0
	for _,x :=  range slx{
		sum +=x
	}
	return sum,"Both returns"
}


func me(){
	defer func(){
		fmt.Println("I will run after my parent method invoked")
	}()
	fmt.Println("I will run First")
}

type person struct{
	first,last string
	age int
}
func (p person) speak(){
	fmt.Println("Person name ",p.first ," & Age is ",p.age)
}
