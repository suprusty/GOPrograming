package main

import "fmt"

func main(){
	arrayExp()
	fmt.Println("-----")
	//sliceExp()
}

var arraynames [5]int
var arrayX = [10]int{3,4,5,6}
func arrayExp(){
	//schools[2] := {"KBJ","BLP"}

	fmt.Println(arraynames)
	fmt.Println(arrayX)

	var x [2]int
	fmt.Println(x)

	x[1] = 99
	fmt.Println(x)
	for i,v := range x {
		fmt.Println(i,v)
	}
	fmt.Println(x[0:1])
}

