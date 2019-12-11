package main

import "fmt"


func main(){
	fmt.Println("Example 1")
	example()

	fmt.Println("Example 2")
	example2()

	fmt.Println("--DELETE SLICE----")
	deleteSlice()

}
func example(){
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x, 43, 43, 43, 43, 43, 43, 44) // new underlying array allocated
	y[0]=1000
	y[1]=2000
	y[2]=3000
	fmt.Println(x)
	fmt.Println(y)
}

func example2() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
	fmt.Println(y)
}

func deleteSlice(){
	x := []int{4,5,6,7,8}
	fmt.Println(x)
	x = append(x[:2],x[4:]...)
	fmt.Println(x)
}