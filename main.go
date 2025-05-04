package main

import "fmt"

// var (
//   a = 10 
//   b = 40
// )

func call(p,q int, func op(a, b int)) {
	op(p,q)
}

func sum(a, b int) {
	c := a + b
  fmt.Println(c)
}



func main() {
call(6,7, sum)
}
