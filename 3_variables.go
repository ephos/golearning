package main

import (
	"fmt"
)

func main() {
	//declare a string
	var a = "initial"
	fmt.Println(a)

	//declare multiple variables at once
	var b, c int = 7, 9
	fmt.Println(b, c)

	//declare multiple variables at once, let compiler infer the type
	var d, e = 2, 1
	fmt.Println(d, e)

	//declare a boolean
	var f = true
	fmt.Println(f)

	//variables can be declared with no value, they will initialize with a 0 value
	var g int
	fmt.Println("g will be 0: ", g)

	var i float64
	fmt.Println("i will be 0: ", i)

	var j byte
	fmt.Println("j will be 0: ", j)

	var k bool
	fmt.Println("k will be false: ", k)

	var h string
	fmt.Println("h will be empty: ", h)

	//shot hand variable declaration is used with := (I like to call this the "Rick and Morty face")
	l := "hey im a shorthand variable"
	fmt.Println(l)
}
