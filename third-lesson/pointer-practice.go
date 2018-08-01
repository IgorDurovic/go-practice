package main

// GO POINTERS ARE PRETTY SIMILAR TO MOST POPULAR LANGUAGES
// THAT SUPPORT POINTERS SO NOTHING TOO CONFUSING HERE

import ("fmt")

func main() {
	x := 15
	
	// assigning the memory address of variable x
	// to a variable a
	a := &x

	fmt.Println("The value of x=", x, "and the memory address of x is", a)

	// basic dereferencing works the same way as with c++
	fmt.Println("Deferencing the reference to x:", *a)

	// to change the value at the memory address of x use the
	*a = 2357
	fmt.Println("x was assigned", x, "using a pointer to its memory address")


}