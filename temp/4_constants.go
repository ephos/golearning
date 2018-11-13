package main

import (
	"fmt"
	"math"
	"reflect"
)

//constant statements can go anywhere a variable can.
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	fmt.Println(reflect.TypeOf(n))

	//constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	//numeric constants have no type until given one by an cast
	fmt.Println(int64(d))
	fmt.Println(reflect.TypeOf(int64(d)))

	fmt.Println(math.Sin(n))
	fmt.Println(reflect.TypeOf(math.Sin(n)))
}
