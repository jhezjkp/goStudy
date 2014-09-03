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
    fmt.Printf("%d\n", 100)
    fmt.Printf("%d\n", 0x100)
    fmt.Printf("%x\n", 0x100)
    fmt.Printf("%X\n", 0x100)
    fmt.Printf("%X\n", 256)
    fmt.Printf("%5d\n", 256)

    fmt.Println("===运算符===")
    fmt.Printf("%d\n", 2 ^ 3)
    fmt.Printf("%d\n", ^ 2)
}
