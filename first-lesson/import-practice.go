package main

// WHENEVER YOU NEED INFORMATION ABOUT A FUNCTION
// IN A GO PACKAGE USE 
//      godoc [package path] [function name]
// IN TERMINAL

// import options:

// OLD WAY
// import "fmt"
// import "math"

// GO WAY
import ("fmt"
		"math")

// Can also have packages within packages.
// If you want to import math subpackage
// rand then
import "math/rand" // USE SLASH INSTEAD OF DOT OPERATOR

func foo () {
	fmt.Println("Result of function call")
	fmt.Println("Using math package: square root of 666 is", math.Sqrt(666))
}

// use math package
func main() {
	// Every externally accessibly function will begin with a
	// capital letter
	//
	// if P is capitalized that function will be exported by go
	// so:
	fmt.Println("Always use capital first letter when calling external functions")
	// instead of:
	// fmt.printLn("This won't work")

	// standard way of calling functions
	foo()

	// without source of entropy (seeding) go rand is completely
	// deterministic
	fmt.Println("A number between 1-100", rand.Intn(100))

}