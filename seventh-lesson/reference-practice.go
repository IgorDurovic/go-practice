package main

// TWO TYPES OF METHODS:
// 		VALUE RECEIVERS - PASS "THIS" BY VALUE
//		POINTER RECEIVERS - PASS "THIS" BY REFERENCE
//
// THIS LESSON WILL DEMONSTRATE POINTER RECEIVERS

import ("fmt")

// same struct from previous lesson
//
// syntax for creating struct. note that the keyword type
// indicates that defining a struct is also defining a
// type
type car struct {
	// 16 bit unsigned int
	gas_pedal uint16
	brake_pedal uint16

	// 16 bit signed in
	steering_wheel int16

	// double precision floating point
	top_speed_kmh float64
}

// This is a method belonging to type car. This 
// time c is a reference to the instance that's
// doing the calling.
//
// NOTE: in this case, c is analogous to the 
// special "this" object from other popular
// languages
func (c *car) set_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed;
}

func main() {
	// you can also assign an instance of the car
	// using c++ style initializer lists
	car_one := car {1, 2, 3, 4}

	fmt.Println("The old top speed of our car is", car_one.top_speed_kmh)

	// set a new top speed by calling set_top_speed
	// using our car object
	car_one.set_top_speed(999)

	fmt.Println("The new top speed is", car_one.top_speed_kmh)

}