package main

import "fmt"

func main(){
foo(1,2,3,4,5,6)
}
func foo(x ...int){
	for i,v := range x{
		fmt.Println("Index ",i ,v )
	}
}
