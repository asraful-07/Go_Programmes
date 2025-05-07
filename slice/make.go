package main

import "fmt"

func main() {
 s := make([]int, 6, 7) // len: 6 || cap: 7
 s := make([]int, 3)

 s[0] =3 

 fmt.Println(s)
 fmt.Println(len(s))
 fmt.Println(cap(s))
}
