package main

import "fmt"

var (
  a = 10 
  b = 40
)

func sum(a, b int) { //parameter a, b
	c := a + b
  fmt.Println(c)
}

func main() {
sum(6, 7) //argument 6, 7
	
}
