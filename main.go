package main

import "fmt"



func main() {
x := 20
p := &x // p ভেরিয়েবল, x এর এড্রেস রাখে

*p = 30 // এড্রেস p তে থাকা মান (x) পরিবর্তন করে

fmt.Println(x)  // 30
fmt.Println(p)  // x এর এড্রেস
fmt.Println(*p) // 30 (এড্রেসে থাকা মান)

}
