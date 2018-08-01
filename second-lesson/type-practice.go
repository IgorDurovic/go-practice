package main

import ("fmt"
		"math")

// go is a statically typed language (thank god)

// paramter types as parameters and return values.
// types are post fixed for both variables and functions.
// go has useful shorthand for typing multiple variables
func add(x float64, y float64) float64 {
	return x + y
}

// functions can return multiple values
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	// variable declaration syntax
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(add(num1, num2))

	// declaring variables with shorthand + inferred types.
	// var keyword isn't needed within functions. This shorthand
	// infers a type of float64 so it cannot be used as anything
	// else without coercion
	//
	// := is shorthand for declaration + assignment
	a, b := 3.14, 4.13

	// the go compiler doesn't allow for defining variables
	// without using them so I'll add another print call to
	// allow my program to compile
	fmt.Println("The sqare root of a + b is", math.Sqrt(add(a, b)))

	// constant variables work like they do with most languages.
	// value must be assigned at definition. Type inference doesn't
	// work with := because constants must be assigned at declaration
	const c = 666
	fmt.Println("Printing out a const", c)

	// demonstrating multiple return values
	w1, w2 := "Hello", "World!"
	fmt.Println(multiple(w1, w2))

	// type conversion
	var x int = 112358
	var y float64 = float64(x)
	fmt.Println("Demonstrating type converstion: int::x + float64::y =", add(float64(x), y))
}