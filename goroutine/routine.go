package main

import (
    "fmt"
    "time"
)

var a = 10
const b = 11

func printHello(num int) {
    fmt.Println("Hello Rahat", num)
}

func main() {
    go printHello(1)
    go printHello(2)
    go printHello(3)
    go printHello(4)
    go printHello(5)
    go printHello(6)

    fmt.Println(a, "", b)

    time.Sleep(1 * time.Second) 
}
