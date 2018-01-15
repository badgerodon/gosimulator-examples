package main 

import "fmt"

func main() {
    var x string
    fmt.Println("Enter your name:")
    fmt.Scan(&x)
    fmt.Println("Hello", x)
}
