package main

import (
	"bytes"
	"fmt"
	"reflect"
)
func main() {
	fmt.Println("-------BasicCompare----------")
	basicCompare()
	fmt.Println("-------ByteCompare----------")
	byteCompare()
	fmt.Println("-------ObjectComapre----------")
	objectComapre()
}
func basicCompare() {
	var a []int = nil
	var b []int = make([]int, 0)
	var c []int = make([]int, 0)

	//fmt.Println(a == b) //compiler error: invalid operation: a == b (slice can only be compared to nil)

	fmt.Println(reflect.DeepEqual(a, b)) // false
	fmt.Println(reflect.DeepEqual(b, c)) // true

	b = append(b, 1)
	c = append(c, 1)
	fmt.Println(reflect.DeepEqual(b, c)) // true

	x := 1
	y := 1
	b = append(b, x)
	c = append(c, y)
	fmt.Println(reflect.DeepEqual(b, c)) // true

	b = append(b, 2)
	c = append(c, 3)
	fmt.Println(reflect.DeepEqual(b, c)) // false
}

func byteCompare() {
	var a []byte = nil
	var b []byte = make([]byte, 0)
	var c []byte = make([]byte, 0)

	//bytes.Equal(): Equal returns a boolean reporting whether
	//   a and b are the same length and contain the same bytes.
	//   A nil argument is equivalent to an empty slice.


	fmt.Println(reflect.DeepEqual(a, b)) // false
	fmt.Println(bytes.Equal(a,b)) // true
	fmt.Println(reflect.DeepEqual(b, c)) // true
	fmt.Println(bytes.Equal(b,c)) //true

	b = append(b, '1')
	c = append(c, '1')
	fmt.Println(reflect.DeepEqual(b, c)) // true
	fmt.Println(bytes.Equal(b,c)) // true

	var x byte = '1'
	var y byte = '1'
	b = append(b, x)
	c = append(c, y)
	fmt.Println(reflect.DeepEqual(b, c)) // true
	fmt.Println(bytes.Equal(b,c)) // true

	b = append(b, '2')
	c = append(c, '3')
	fmt.Println(reflect.DeepEqual(b, c)) // false
	fmt.Println(bytes.Equal(b,c)) // false
}

func objectComapre() {
	map1 := map[string]int{
	"x": 1,
	"y": 2,
	}
	map2 := map[string]int{
	"x": 1,
	"y": 2,
	}
	map3 := map[string]int{
	"x": 1,
	"y": 2,
	"z": 3,
	}
	fmt.Println("Map equal: ", reflect.DeepEqual(map1, map2)) // true
	fmt.Println("Map equal: ", reflect.DeepEqual(map1, map3)) // false
}