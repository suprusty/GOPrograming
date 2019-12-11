package main

import "fmt"

type user struct{

}

func main(){
	u := user{}
	sf,val := u.syntextOffunc("sushil",1000)
	fmt.Println(sf,val)
}

func (u user) syntextOffunc(name string,salary int) (string,int){
	fmt.Println("INSIDE Funcation")
	return name,salary
}
