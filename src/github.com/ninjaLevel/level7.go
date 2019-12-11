package main

import "fmt"

type personL struct {
	age int
	name string
}

func main(){
	pl := personL{
		age:34,
		name:"sushil",
	}
	fmt.Println(pl)
	fmt.Println(&pl)
	changeMe(&pl)
	fmt.Println(pl)
	fmt.Println(&pl)
}

func changeMe(p *personL){
	(*p).age = 100
	(*p).name = "I changed"
	//p.name = "I changed or not 	"
}