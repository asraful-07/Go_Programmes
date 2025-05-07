package main

import "fmt"

func  variadic(numbers...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	
}

func main() {
	variadic(1,2,3,4,5,6)
}
