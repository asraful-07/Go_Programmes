package main

import "fmt"

// Adds two numbers and returns the result
func add(num1 int, num2 int) int {
	sum := num1 + num2
	fmt.Println("Sum from add():", sum)
	return sum
}

// Returns sum and multiplication of two numbers
func numbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

// Returns sum, multiplication, and whether the sum is odd (1) or even (0)
func odd(num1 int, num2 int) (int, int, int) {
	sum := num1 + num2
	mul := num1 * num2
	isOdd := sum % 2
	return sum, mul, isOdd
}

func main() {
	a := 10
	b := 20

	// Get sum and product
	p, q := numbers(a, b)

	// Pass those values into 'odd' to get new results
	s, r, oddCheck := odd(p, q)

	fmt.Println("Sum:", s)
	fmt.Println("Product:", r)
	if oddCheck == 1 {
		fmt.Println("The sum is odd")
	} else {
		fmt.Println("The sum is even")
	}
}
