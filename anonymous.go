package main

import "fmt"

var (
	a = 10
	b = 30
)

func main() {
	fun(a int, b int){
		res := a + b 
		fmt.Println(res)
	}(4, 7) //(4, 7) this IIFE
}

func init() {
	fmt.Println("but i am fist run and next you")
}