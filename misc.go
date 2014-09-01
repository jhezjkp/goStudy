package main

import "fmt"

func sayHi(name string) string {
    return "hello, " + name + "!"
}

func init() {
    fmt.Println("function init will be executed first.")
}

func main() {
    fmt.Println(sayHi("vivia"))
    const PI = 3.14
    fmt.Printf("PI=%v\n", PI)
}
