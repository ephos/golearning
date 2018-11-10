package main

import (
	"fmt"
)

func main() {
	//an array that only holds 5 integers, because we just declared the int with no values, all 5 items in the array will get the 'zero value' for int of 0.
	var a [5]int
	fmt.Println("show alpha set:", a)

	//you can set a value by referencing its index.  As always, indexes start at 0, so 4 is the last (5th) element in this array
	a[4] = 100
	fmt.Println("show alpha set:", a)
	fmt.Println("show alpha item [4]", a[4])
	//builtin len function/method will show the length of the array
	fmt.Println("show alpha length:", len(a))

	//an array can also be initialized and assigned values with short hand variable declaration
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("show beta set:", b)

	//build a multi-dimensional array
	var c [2][3]int
	//while i is less than 2, increment 'i' by 1
	for i := 0; i < 2; i++ {
		///while j is less than 3, increment 'j' by 1
		for j := 0; j < 3; j++ {
			//c dimension 1 is equal to i + j
			//c dimension 2 is equal to i + j
			//in iteration 1: since i and j both start as 0, in both arrays (in each dimension) the sum is 0
			//in iteration 2: since i is 1, and j is 1
			c[i][j] = i + j
			fmt.Println("show charlie sets: ", c)
		}
	}
}
