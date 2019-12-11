package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice1(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)]
	slice2 := slice1[2:4]
	inspectSlice1(slice2)

	fmt.Println("*************************")

	// Change the value of the index 0 of slice2.
	slice2[0] = "CHANGED"

	// Display the change across all existing slices.
	inspectSlice1(slice1)
	inspectSlice1(slice2)
}

// inspectSlice exposes the slice header for review.
func inspectSlice1(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			v)
	}
}
