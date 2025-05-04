package main

import "fmt"

func main() {
	// taka := 10 
	const salary = 40000;
	const isBcs = true;
	const hight = 68;
      
	if salary > 10000 && isBcs == true  && hight > 64 {
		fmt.Println("ami raji")
	} else if salary < 20000 || hight > 60 {
		fmt.Println("I am also raji")
	} else {
		fmt.Println("of oooooo")
	}
	
}

// *************

fmt.Println("welcome to my application")

// Get user name input 
var name string
  fmt.Println("Enter your--")
  fmt.Scanln(&name)


// *************

  package main

import (
	"fmt"
)

var (
  a = 20 
  b = 40
)

func sumResult(num int) {
  fmt.Println(num)
}

func add(x int, y int) {
  sum := x + y
  sumResult(sum)

}


func main() {
 add(a, b)
}

// ***************** shadding func
package main

import (
	"fmt"
)

var (
  a = 20 
  b = 40
)

func main() {
 age :=  20

 if age >= 18 {
  a := 12
  fmt.Println("hi i am", a)
 }else{
  fmt.Println("i am ", a)
 }
 
}


// ********************
package main

import (
	"fmt"
  // "example.com/mathlib"
)

var (
  a = 10 
  b = 40
)

func main() {
  fmt.Println("welcome to my application", a)
}

func init() {
  fmt.Println("this init func", a)
  a = 30
}


// **********
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






