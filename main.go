package main

import "fmt"

func add() (result int) {
	fmt.Println("first", result)

 show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5 
	fmt.Println("second", result)

	return result
}

func main() {
   add(4,8)


}
