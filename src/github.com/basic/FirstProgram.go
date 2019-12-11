package main

import "fmt"
var amount int =10000
var name string
 //comp := "Cisco"
const (
	VALID = 10001
	STS = "USA"
)

const (
	a = iota
	b
	c
)
func main()  {
	var localtomain = "localvariable to main"

	age := 100
	fmt.Println("My First commit",localtomain,age,amount,name)
	fmt.Printf("%d\t%#x\t%b",age,age,age)

	x := 10
	fmt.Printf("%d\t%#x\t%b \n",x,x,x)
	y:= 10<<1
	fmt.Printf("%d\t%#x\t%b \n ",y,y,y)


	fmt.Println(a ,b ,c )
	forexp()
	switchexp()
}

func forexp(){

	i :=0
	for{

		if i>10 {
			break
		}
		i++;

		fmt.Println(i)
	}
	fmt.Println("Done !!!!")

	for j := 33; j<=122; j++{
		fmt.Printf("%v\t%#x\t%#U \n",j,j,j)
	}
}

func switchexp(){
	switch {
		case (2 == 4):
			fmt.Println("This should not print")
		case (2==2):
			fmt.Println("This should  print")
			fallthrough;
		case (4==4):
			fmt.Println("This should  print")
		default :
			fmt.Println("Default case")
	}

	switch {
	case true:
		fmt.Println("This should not print")
	case false:
		fmt.Println("This should  print")
	}

	switch "sushil"{
	case "hari":
		fmt.Println("hari")
	case "krish":
		fmt.Println("krish")
		fallthrough;
	case "kumar","prusty,","sushil":
		fmt.Println("sushil")
	default :
		fmt.Println("Default case")
	}
}
