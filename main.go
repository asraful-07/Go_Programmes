package main

import "fmt"
 
func main() {
   x := 10
   y := 12

   increment := func() int{
   x += 1
   return x
   }

   
   decrement := func() int {
      y -= 1
      return y
   }

   fmt.Println(increment())
   fmt.Println(increment())

   fmt.Println(decrement())
   fmt.Println(decrement())
}