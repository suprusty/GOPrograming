package main

import "fmt"

func main(){
	basicMap()
}

func basicMap(){
	//Maps Initializations
	var m1 map[string]int // nil
	var m1a = map[string]int{}
	var m2 = make(map[string]int)
	var m3 = make(map[string]int, 100) // capacity hint but can outgrow

	fmt.Println(len(m1), len(m1a), len(m2), len(m3)) // 0, 0, 0, 0
	fmt.Println(m1, m1a, m2, m3)                     // map[], map[], map[], map[]
	fmt.Printf("%p %p %p %p\n", m1, m1a, m2, m3)     // 0x0, 0x43e260, 0x43e280, 0x43e2a0

	//m1["a"] = 1 // panic: assignment to entry in nil map
	m1a["a"] = 1
	m2["a"] = 1
	m3["a"] = 1


	userMap := make(map[int]int)
	fmt.Println("Null Map :",userMap)
	userMap[1] = 10
	userMap[2] = 20
	userMap[5] = 50
	fmt.Println("Value of Map :",userMap)

	v1 := userMap[5]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(userMap))

	delete(userMap,2)
	fmt.Println("After Deletion :", userMap)

	_, prs := userMap[5]
	fmt.Println("prs:", prs)

	userInstialiseMap := map[int]string{1:"One",2:"two",3:"three"}
	fmt.Println("Intialised Map :",userInstialiseMap)


	advanceMap()
}
type user struct {
	name    string
	surname string
}
var globaluserMap = map[string]user{"Roy":user{"Rob", "Roy"},"Ford":user{"Henry", "Ford"},"Mouse":{"Mickey", "Mouse"},
	"Jackson":{"Michael", "Jackson"}}
func advanceMap(){
	fmt.Println(globaluserMap)

	for key, value := range globaluserMap {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map and notice the
	// results are different.
	for key := range globaluserMap {
		fmt.Println(key)
	}
}
