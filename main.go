package main

import "fmt"

func callback(name string) {
	fmt.Println("Hi brother, ", name)
}

func greet(sayHi func(string)) {
    sayHi("Rahat")
}

func main() {
 greet(callback)
}