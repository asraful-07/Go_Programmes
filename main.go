package main

import "fmt"

func main() {
 fmt.Println("what is This i am first execute")
  x := func(a,b int) {
	ab := a + b
    fmt.Printf("%d\n", ab)
	}
  x(2,2)

defer fmt.Println("This defer func")	
}
