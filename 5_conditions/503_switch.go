package main

import (
	"fmt"
)

func main() {
	day := 6

	// single case
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default: // default is when any case is not matched
		fmt.Println("Not a weekday")
	}

	// multi case
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// Program to print the day of the week using fallthrough in switch
	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		fallthrough

	case 4:
		fmt.Println("Wednesday") // case 4 also gets executed if case 3 due to fallthrough keyword

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")
	}

	// switch without expression
	var numberOfDays = 28
	switch {

	case 28 == numberOfDays:
		fmt.Println("It's February")

	default:
		fmt.Println("Not February")
	}

}
