package main

import (
	"fmt"
  // "example.com/mathlib"
)

var (
  a = 10 
  b = 40
)

// func add(p, q int) {
//   res := p + q 
//   fmt.Println(res)
// }


func main() {
  // we invoke the add func here
  // add(7, 9) //invoke
  // anonymous func 
  // Immediately Invoked fun expression

  // a := 10 //expression
  // // if expression
  // if a > 0 {
  //   fmt.Println("this is a")
  // }

func(a int, b int) { //if block
c := a + b 
fmt.Println(c)
}(6, 7)

  // fmt.Println("welcome to my application", a)
}

func init() {
  fmt.Println("this init func")
  // a = 30
}

