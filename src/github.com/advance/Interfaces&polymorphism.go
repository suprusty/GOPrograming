package main

import "fmt"
//A value can be more then one type
//Incase marker interface will be any value type
type person struct {
	first string
	last  string
}

type secretAgent struct{
	person
	ltk bool
}

func (p person) speak(){
	fmt.Println("Persion -", p.first , " , " , p.last)
}
func (sa secretAgent) speak(){
	fmt.Println("secretAgent -", sa.first , " , " , sa.ltk)
	fmt.Println("secretAgent having persion -", sa.person)
}


type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h.speak)
}

func main(){
		p := person{
			first : "MD",
			last :"Prusty",
		}
		sc := secretAgent{
			person:person{
				first : "sushil",
				last :"Prusty",
			},
			ltk:true,
		}
		sc1 := secretAgent{
			person:person{
				first : "Shiv",
				last :"Prusty",
			},
			ltk:false,
		}
	bar(p)
	bar(sc)
	bar(sc1)
}