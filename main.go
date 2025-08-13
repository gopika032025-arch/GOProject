package main

import (
	"MYPROJECT/simplecalc"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	a, b := 7, 2
	fmt.Println(simplecalc.Add(a, b))

	c, d := 15, 3
	fmt.Println(simplecalc.Sub(c, d))

	e, f := 30, 4
	fmt.Println(simplecalc.Div(e, f))

	x, y := 20, 7
	fmt.Println(simplecalc.Mul(x, y))

	// --- if / else if / else ---
	num := 26
	if num > 26 {
		fmt.Println("Number is greater than 26")
	} else if num == 26 {
		fmt.Println("Number is exactly 26")
	} else {
		fmt.Println("Number is less than 26")
	}

	// --- switch case ---
	day := "Wednesday"
	switch day {
	case "Wednesday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("It's another day")
	}

	// --- for loop ---
	fmt.Println("Counting from 1 to 9:")
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}
	// --- pointers ---
	j := 9
	fmt.Println("Before:", j)
	increase(&j)
	fmt.Println("After:", j)

	// --- defer  ---

	defer fmt.Println("Second")
	defer fmt.Println("Middle")
	defer fmt.Println("First")

}

func increase(num *int) {
	*num = *num + 1
}
