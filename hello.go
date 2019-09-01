package main

import "fmt"

// Hello retruns the strig "Hello, world"
func Hello(name string) string {
    return "Hello, " + name
}

func main() {
    fmt.Println(Hello("world"))
}