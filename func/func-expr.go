package main

import "fmt"

var (
  a = 10 
  b = 40
)


func main() {

  add := func(a, b int){
		res := a + b 
		fmt.Println(res)
	}
	add(4, 7)
	
}

func init() {
	fmt.Println("but i am fist run and next you")
}