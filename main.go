package main

import "fmt"

func first() (result int) {
	fmt.Println("Result", result)

	defer func() {
		result += 2
	}()

	result = 6
	fmt.Println("defer", result)

	return // Returns 6 + 2 = 8
}

func first1() int{
   result := 0
	fmt.Println("Result", result)

	defer func() {
		result += 2
	}()

	result = 6
	fmt.Println("defer", result)

	return result // Returns 6 + 2 = 8
}

func main() {
  a := first()
  fmt.Println("main a", a) // 8

  b := first1()
  fmt.Println("main b", b) // 6
}
