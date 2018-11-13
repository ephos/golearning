package main

import (
	"fmt"
)

func main() {
	//create an empty map
	//[key type]value type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
}
