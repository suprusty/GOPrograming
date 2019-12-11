package main

import "fmt"

var globalA = [5]int{1,2,3,4,5}
var globals = []int{12,13,14}
func main(){
	fmt.Println(globalA)

	localArray := [...]int{4,5,6,7,8}
	fmt.Println(localArray)

	for i,v := range localArray{
		fmt.Println(i,v)
	}

	localS := []int{}
	localS = append(localS,100)
	localS = append(localS,101)

	for i,v := range localS{
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n",localS)

	sliceofSlice()
}

func sliceofSlice(){
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xs3 := []string{"kbj", "sbp"}
	xs4 := []string{"ODI", "BANG", "VSKP"}
	xs5 := []string{"CISC", "IBM"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2,xs3,xs4,xs5}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

func expMap(){
	mexp := make(map[int][]string)
	mexp[0] =[]string{"James", "Bond", "Shaken, not stirred"}

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

