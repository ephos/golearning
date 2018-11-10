package main

import (
	"fmt"
)

func main() {
	//concatenate a string
	fmt.Println("go" + "lang")

	//expressions printed
	fmt.Println("1+1 =", 1+1)         //integers
	fmt.Println("7.0/3.0 =", 7.0/3.0) //floats

	//booleans
	fmt.Println(true && false) //true and false = false
	fmt.Println(true || false) //true or false = true
	fmt.Println(!true)         //not true
}
