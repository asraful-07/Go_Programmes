package main

import "fmt"

// Struct definition with a type name
type Person struct {
	name   string
	age    float32
	class  [4]int8
	isBool bool
}

func main() {
	// Declare a variable of type Person and initialize it
	Person1 := Person{
		name:   "Rahat",
		age:    22.1,
		class:  [4]int8{1, 2, 3, 4},
		isBool: true,
	}

	// Print values
	fmt.Printf("Name: %s\n", Person1.name)
	fmt.Printf("Age: %.2f\n", Person1.age)
	fmt.Printf("Class: %v\n", Person1.class)
	fmt.Printf("isBool: %v\n", Person1.isBool)
}
