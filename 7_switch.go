package main

import (
	"fmt"
	"time"
)

func main() {
	//basic switch, not too dissimilar to other langauges.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//like other languages default case will always happen if no others matched.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	//again, similar to .Net a switch without an expression will roughly compile down to an if/else statement/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon.")
	}

	//a more advanced example to compare type
	//whatAmI is equal to a function that takes a single parameter input 'i' as a type of 'interface'.
	//the function contains a switch statement that will evaluate the type of the passed entity/object/struct.
	whatAmI := func(i interface{}) {
		//switch then evaluates i.(type) as variable 't' to determine what it is.
		//.(type) syntax is only used inside a switch, which is why this doesn't require the reflector library.
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool.")
		case int:
			fmt.Println("I am an int.")
		default:
			fmt.Printf("I don't know what type %T is!", t)
		}
	}
	//using this function we canb
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
