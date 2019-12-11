package main

import "fmt"

func passarg(x [3]int){
	x[0] = 1000;
	fmt.Println(x)
}

func passslic(y []int){
	y[0] = 1000;
	fmt.Println(y)
}

func pfa(pa *[3]int) {
	pa[0] = 1000
}

func main(){
	arr1 := [3]int{1,2,3}
	passarg(arr1);
	fmt.Println("After method invoked array value : -",arr1)
	sli1 := []int{1,2,3}
	passslic(sli1)
	fmt.Println("After method invoked Slice value : -",sli1)


	pfa(&arr1)
	fmt.Println("After a as pointer:", arr1, sli1)
}
