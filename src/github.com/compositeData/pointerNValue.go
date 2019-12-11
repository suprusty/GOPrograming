package main

import (
	"fmt"
	"math"
)

type vertix struct{
	X,Y float64
}

func (v vertix) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v vertix) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
func (vp *vertix) IneffectiveScale(f float64){
	vp.X = vp.X * f
	vp.Y = vp.Y * f
}
func main(){
	v := vertix{
		5,5,
	}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println("With out Pointer value ",v)

	v.IneffectiveScale(2)
	fmt.Println("With Pointer value ",v)

}
