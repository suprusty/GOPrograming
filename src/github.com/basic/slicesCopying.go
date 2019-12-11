package main

import "fmt"

func main() {
	//n = copy(dst, src)

	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}

	fmt.Println("\nwhen dst is nil: ")
	fmt.Printf("\tbefore copy: dst=%v, src=%v\n", s1, s2)
	n1 := copy(s1, s2)
	fmt.Printf("\tafter copy: n=%d, dst=%v, src=%v\n", n1, s1, s2)
	fmt.Println("\tdst == nil", s1 == nil)

	fmt.Println("\nwhen dst is smaller: ")
	fmt.Printf("\tbefore copy: dst=%v, src=%v\n", s2, s3)
	n2 := copy(s2, s3)
	fmt.Printf("\tafter copy: n=%d, dst=%v, src=%v\n", n2, s2, s3)

	fmt.Println("\nwhen dst is bigger: ")
	fmt.Printf("\tbefore copy: dst=%v, src=%v\n", s3, s4)
	n3 := copy(s3, s4)
	fmt.Printf("\tafter copy: n=%d, dst=%v, src=%v\n", n3, s3, s4)
}