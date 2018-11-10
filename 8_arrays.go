package main

import (
	"fmt"
)

func main() {
	//an array that only holds 5 integers, because we just declared the int with no values, all 5 items in the array will get the 'zero value' for int of 0.
	var a [5]int
	fmt.Println("emp:", a)

	//you can set a value by referencing its index.  As always, indexes start at 0, so 4 is the last (5th) element in this array
	a[4] = 100
	fmt.Println("show set:", a)
	fmt.Println("show item [4]", a[4])
	//builtin len function/method will show the length of the array
	fmt.Println("show length:", len(a))
}
