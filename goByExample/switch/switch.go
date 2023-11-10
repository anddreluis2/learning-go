package main
import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as ")
	switch i {
		case 1: 
			fmt.Println("one")
		case 2: 
			fmt.Println("two")
		case 3: 
			fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend!")
	default:
		fmt.Println("it's week-day!")
	}

	t:= time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("im a boolean")
			case int:
				fmt.Println("im a integer")
			case string:
				fmt.Println("im a string")
			default:
				fmt.Println("don't know the type", t)
		};
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("loool")
}