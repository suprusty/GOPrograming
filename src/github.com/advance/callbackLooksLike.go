package main

import "fmt"
//passing a func as an argument
func main(){
	math(2,5,add)
	math(2,5,mult)
}

func add(x,y int){
	fmt.Println(x+y)
}
func mult(x,y int){
	fmt.Println(x*y)
}


func math(x int,y int,f func(a,b int)){
	f(x,y);
}