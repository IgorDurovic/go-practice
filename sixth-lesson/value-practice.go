package main

// TWO TYPES OF METHODS:
// 		VALUE RECEIVERS - PASS "THIS" BY VALUE
//		POINTER RECEIVERS - PASS "THIS" BY REFERENCE
//
// THIS LESSON WILL DEMONSTRATE VALUE RECEIVERS

import ("fmt")

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

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

// our first method. Belongs to type car
// where c is the value of the instance being
// used.
//
// NOTE: don't confuse c with the "this" special
// object in other languages. c does not hold a 
// reference to the car object doing the calling
func (c car) kmh() float64 {
	// if you tried to changed the value of a member variable
	// of c it will only change the local copy
	//
	// ex. c.top_speed_kmh = 10
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func main() {
	// you can also assign an instance of the car
	// using c++ style initializer lists
	car_one := car {1, 2, 3, 4}

	fmt.Println("Gas pedal intensity:", car_one.gas_pedal)
	
	// call the kmh() member method of car_one
	fmt.Println("The car's current speed in kmh:", car_one.kmh())

	fmt.Println("The car's current speed in mph:", car_one.mph())


}