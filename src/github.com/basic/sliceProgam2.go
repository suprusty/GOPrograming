package main

import "fmt"

var slicenames []int
var slicey = []int{3,4,5,6}

func main(){
	sliceExp()
}

func sliceExp(){
	fmt.Println(slicenames)
	fmt.Println(slicey)


	x := []int{4,5,6,7,8}
	for i,v := range x {
		fmt.Println(i,v)
	}
		fmt.Println("--slice Explansation ---")
	sliceExplansation()
	fmt.Println("--slice Appending ---")
	sliceAppending()
}

func sliceUsing_Make(){
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
			x[0] = 42
			x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multidimensionalslice(){
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}


func sliceExplansation(){
		s := make([]string, 3)
		fmt.Println("emp:", s)

		// We can set and get just like with arrays.
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])

		// `len` returns the length of the slice as expected.
		fmt.Println("len:", len(s))

		// In addition to these basic operations, slices
		// support several more that make them richer than
		// arrays. One is the builtin `append`, which
		// returns a slice containing one or more new values.
		// Note that we need to accept a return value from
		// append as we may get a new slice value.
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

		// Slices can also be `copy`'d. Here we create an
		// empty slice `c` of the same length as `s` and copy
		// into `c` from `s`.
		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy:", c)

		// Slices support a "slice" operator with the syntax
		// `slice[low:high]`. For example, this gets a slice
		// of the elements `s[2]`, `s[3]`, and `s[4]`.
		l := s[2:5]
		fmt.Println("sl1:", l)

		// This slices up to (but excluding) `s[5]`.
		l = s[:5]
		fmt.Println("sl2:", l)

		// And this slices up from (and including) `s[2]`.
		l = s[2:]
		fmt.Println("sl3:", l)

		// We can declare and initialize a variable for slice
		// in a single line as well.
		t := []string{"g", "h", "i"}
		fmt.Println("dcl:", t)

		// Slices can be composed into multi-dimensional data
		// structures. The length of the inner slices can
		// vary, unlike with multi-dimensional arrays.
		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)

}

func sliceAppending(){
		// Create a slice with a length of 5 elements.
		slice := make([]string, 5)
		slice[0] = "Apple"
		slice[1] = "Orange"
		slice[2] = "Banana"
		slice[3] = "Grape"
		slice[4] = "Plum"

		// You can't access an index of a slice beyond its length.
		//slice[5] = "Runtime error"

		// Error: panic: runtime error: index out of range

		fmt.Println(slice)
		fmt.Println("---- Before value pass ---")
		fmt.Printf("%p \n",&slice)
		for i, v := range slice {
			fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			v)
		}

		inspectSlice(slice)

}

func inspectSlice(slice []string) {
	fmt.Println("---- After value pass ---")
	fmt.Printf("%p \n",&slice)
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))

	for i, v := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			v)
	}
}
