package main

import (
	"MYPROJECT/simplecalc"
	"fmt"
)

// -----
type Car struct {
	Name  string
	Speed int // in km/h
}

// ---Method increase speed
func (c *Car) accelerate(extra int) {
	c.Speed = c.Speed + extra
}

//----function to add bonus marks-----

func addBonus(marks *int) {
	*marks = *marks + 5
}

// --- Declaration---Structures -----
type College struct {
	Name  string
	Count int
}

func main() {
	// hello world

	fmt.Println("Hello World")

	//----Slice-Declaration----
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// int......

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
	j := 8
	fmt.Println("Before:", j)
	increase(&j)
	fmt.Println("After:", j)

	// --- defer  ---

	defer fmt.Println("Middle")
	defer fmt.Println("First")
	defer fmt.Println("Second")
	//---------------//-------------------//
	// ------Structure call---- --- Structure Example ----

	college := College{
		Name:  "NIU",
		Count: 4000,
	}
	fmt.Println(college.Name)
	fmt.Println(college.Count)
	fmt.Println(college)

	//--- Arrays -- Fixed length seq of num ---

	//Array--INT---

	var A [7]int
	fmt.Println(A) //Default values = 0//

	A[4] = 8
	fmt.Println(A)

	//Array--STRING-

	B := [5]string{"Northern", "Illinois", "University", "NIU", "Chicago"}

	fmt.Println(B)

	//--Array--Range-----

	for i, v := range B {
		fmt.Println(i, v)
	}

	//---Slices---
	//----len & cap------

	m := []int{1, 2, 3}
	fmt.Println(len(m))
	fmt.Println(cap(m))
	fmt.Println(m)

	g := [6]int{3, 4, 5, 6, 7, 8}
	slice = g[2:4] ////Slice from index 2 to 3
	fmt.Println(slice)
	fmt.Println(g)

	n := [5]int{3, 4, 5, 6, 7}
	slice = n[:4] //Slice from beginning till index 3
	fmt.Println(slice)
	fmt.Println(n)

	o := [9]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = o[4:] //Slice from index 4 till end
	fmt.Println(slice)
	fmt.Println(o)

	//-----append-----
	s := make([]int, 4, 8) //// slice with length=4, capacity=8
	s = append(s, 4)       // Add element at the end
	s = append(s, 3)
	fmt.Println(s)

	//-----Class 3 --Maps//

	fruits := map[string]int{
		"Apple":  100,
		"Banana": 40,
	}

	//-----add new fruit----
	fruits["Orange"] = 60

	//----- printing the price of Apple
	fmt.Println("Apple price:", fruits["Apple"])

	//------check if Mango is present
	if val, ok := fruits["Mango"]; ok {
		fmt.Println("Mango price is", val)
	} else {
		fmt.Println("Mango is not available")
	}

	for i, v := range fruits {
		fmt.Println(i, v)

	}

	// For deleting

	delete(fruits, "Apple")

	//--------Pointers & Func

	studentMarks := 40
	fmt.Println("Before Bonus:", studentMarks)

	addBonus(&studentMarks) // passing address
	fmt.Println("After Bonus:", studentMarks)

	//----------Method--------

	myCar := Car{Name: "Toyota", Speed: 60}

	fmt.Println("Before:", myCar.Name, "speed is", myCar.Speed)

	myCar.accelerate(20) // method call

	fmt.Println("After:", myCar.Name, "speed is", myCar.Speed)
}

func increase(num *int) {
	*num = *num + 1
}
