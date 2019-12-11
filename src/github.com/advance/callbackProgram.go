package main

import (
	"fmt"
)

func main(){
	ii := []int{1,2,3,4,5,6,7,8,9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	evensum := even(ii...)
	fmt.Println("evensum numbers", evensum)

	fmt.Println(evenCallBack(sum,ii...))
}

func sum(xi ...int) int{
	fmt.Printf("%T\n", xi)
	sum := 0
	for _,v := range xi{
		sum += v
	}
	return sum
}


func even(ev ...int) int{
	sumEven := 0
	for _,v := range  ev {
		if (v%2 == 0) {
			sumEven += v
		}
	}
	return sumEven
}

func evenCallBack(sumF func(xi ...int) int,ev ...int) int{
	var yi []int
	for _,v := range  ev {
		if (v%2 == 0) {
			yi = append(yi,v)
		}
	}
	fmt.Println(yi)
	return sumF(yi...)
}
