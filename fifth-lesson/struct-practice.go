package main

// GO DOES NOT HAVE ANY CONCEPT OF CLASSES AND INSTEAD USES
// STRUCTS FOR MOST RELATED PURPOSES

import ("fmt")

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

func main() {
	// assigned an instance of the custom struct type
	// using json like syntax
	car_one := car {gas_pedal: 1111,
					brake_pedal: 0,
					steering_wheel: 2222,
					top_speed_kmh: 999}

	// you can also assign an instance of the car
	// using c++ style initializer lists
	car_one = car {1, 2, 3, 4}

	// access the fields of the struct instance
	// the typical way (with the dot operator)
	fmt.Println("The top speed of our car struct instance is", car_one.top_speed_kmh)
}