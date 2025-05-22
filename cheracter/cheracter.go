package main

import "fmt"

func main() {
  var a float32 = 10.2345
  var b string = "MY Name is Rahat"
  var c bool =  false
  var d int8 = 123
  var e rune = 'অ' 
  var j rune = '❤'

  fmt.Printf("%.2f\n", a)
  fmt.Printf("%Ts\n", b)
  fmt.Printf(b)
  fmt.Printf("%v\n", c)
  fmt.Printf("%d\n", d)
  fmt.Printf("%c\n", e)
  fmt.Printf("%c\n", j)
}
