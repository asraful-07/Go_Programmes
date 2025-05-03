// Global Scope
/************  go mod init example.com  ************/
package main

import "fmt"

var (
  a = 20 
  b = 40
)

func add(x int, y int){
   z := x + y 
   fmt.Println(z)
}

func main() {
p := 20
q := 40

add(p, q)

add(a, b)

add(q, a)
}

// Local Scope

package main

import "fmt"

var (
  a = 20 
  b = 40
)

/*

block = { }

*/

func main() {
x := 18

if x >= 18 {
  p := 10
  fmt.Println("your matured boy")
  fmt.Println("you have", p, "friends")
}

}


// Package Scope
package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
  a = 20 
  b = 40
)


func main() {
  fmt.Println("Addition done") 

  mathlib.Add(4, 7)
  
}




