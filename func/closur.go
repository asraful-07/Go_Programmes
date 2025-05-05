package main

import "fmt"

var (
  a = 10 
  b = 40
)

func outer() func() {
	money := 100
	age := 22 

	fmt.Println("my age =", age)

	show := func(){
		money = a + b
		fmt.Println("money", money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}


func main() {
call()
}

func init() {
	fmt.Println("=== Bank ===")
}