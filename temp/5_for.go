package main

import (
	"fmt"
)

//go has for loops like many other langauges, it does not however have a foreach, a for with a range is the equivalent.

func main() {

	//i initializes as 1, loop while j is less than or equal to 3, increment by 1 within the loop
	fmt.Println("Simple 'for' loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//j initializes at 7, loop while j is less than or equal to 9, increment j by 1
	fmt.Println("Classic 'for' loop with index defined in 'for'")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//for loop without a condition is comperable to a do or while loop, it will run until broke out of with 'break'
	for {
		fmt.Println("This will loop forever without a 'break' statement!")
		break
	}

	//for loops can also be broken out of early with 'continue'
	for n := 0; n <= 5; n++ {
		if n == 3 {
			fmt.Println("n is equal to 3, skipping to next iteration of loop")
			continue
		}
		fmt.Println(n)
	}
}
