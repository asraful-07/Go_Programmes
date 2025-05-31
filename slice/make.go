package main

import "fmt"

func main() {
 s := make([]int, 6, 7) // len: 6 || cap: 7
 p := make([]int, 3)  // len: 3 || cap: 3

 s[0] =3 

 fmt.Println(p, len(p), cap(p)) // [0,0,0] len : 3 , cap : 3

 fmt.Println(s)  // s = [3,0,0,0,0,0]
 fmt.Println(len(s)) // len: 6 
 fmt.Println(cap(s))// cap: 7
}
