package main

import ( "fmt" 
        "time"
    )

func HelloPrint(a,b int)  {
	ab := a + b
	fmt.Println("hello", ab)
}


func main() {
	x := 4
 defer fmt.Println("hi how are you", x)

  go HelloPrint(2, 4)
  go HelloPrint(2, 4)
  go HelloPrint(2, 4)
  go HelloPrint(2, 4)
  time.Sleep( 1 * time.Second)
}
