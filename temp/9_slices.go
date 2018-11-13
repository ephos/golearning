package main

import (
	"fmt"
)

func main() {

	//create a slice with a non zero size, this will create a slice that contains 3 strings, the 3 strings will be zero valued
	s := make([]string, 3)
	fmt.Println("empty:", s)

	//set and get is possible like arrays
	s[0] = "wilkenshire"
	s[1] = "winston"
	s[2] = "cromwell"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("length:", len(s))

	//you can append and grow a slice dynamically as it is not fixed like arrays (like .Net)
	//this under the hood is creating a new slice with the original stuff and the new items
	s = append(s, "bob")
	s = append(s, "gary", "mikey")
	fmt.Println("appended:", s)

	//slice can be copied, a slice needs to have an initial size created.
	//in this case rather than defining an integer we use the interger output which is the length of the 's' slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//"low high" AKA 'slicing' syntax.  The number on the left is where the count starts, the index number on the right is where it ends.
	//the number on the right will always be excluded

	//slices support a special slice operator which gets (slices) a subset of the slice
	//for instance in this example it will get elements 2,3,4 from the slice s, it stores this in l
	//don't get mixed up, this means it will start at 2 and end at 5, 5 will not be counted however.
	l := s[2:5]
	fmt.Println("slice subset 's[2:5]' (slices up from s[2],  but excluding s[5]):", l)

	//slice up to (but excluding s[5])
	//note!  since we are reusing 'l' we can't create it as if it is new, therefore we use '=' instead of ':='.  This just assigns the new value.
	l = s[:5]
	fmt.Println("slice subset 's[:5]' (slices up to (but excluding) s[5]):", l)

	//slice up from (and including s[2])
	l = s[2:]
	fmt.Println("slice subset 's[2:]' (slices up from (and including) s[2]):", l)

	//similar to an array, slices can be declared with initial values as well
	t := []string{"g", "h", "i"}
	fmt.Println("populated slice", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two dimensional slice:", twoD)
}
