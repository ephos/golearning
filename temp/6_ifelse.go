package main

import (
	"fmt"
)

//go does not require parenthesis around conditions like in PowerShell or C#

func main() {
	//basic if/else
	if 7%2 == 0 {
		fmt.Println("7 is even!")
	} else {
		fmt.Println("7 is odd!")
	}

	//like .net languages you can have an if without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//a statement can have pre-conditions, any variables declared can be used in all branches
	if num := 10; num < 0 {
		fmt.Println(num, "num is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
