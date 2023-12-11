package main

import "fmt"
    var lines string = "-------------------------"
func main() {
    fmt.Println(lines)
    // Print hello world
    fmt.Println("hello world")
    fmt.Println(lines)
    // Simple calculations
    fmt.Println("1 + 1 = ", 1+1)
    fmt.Println("10 / 5 = ", 10/5)
    fmt.Println("2 * 8 = ", 2*8)
    fmt.Println(lines)
    // For loops
    a := 1
    for a <= 3 {
        fmt.Println("This is number ", a)
        a = a + 1
    }
    fmt.Println(lines)
    // If-else statements
    fmt.Println("Is 7 even or odd?")
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    fmt.Println("Is 8 even or odd?")
    if 8%2 == 0 {
        fmt.Println("8 is even")
    } else {
        fmt.Println("8 is odd")
    }
    // Switch case
    fmt.Println(lines)
    var in string
    fmt.Scanln(&in)
        switch in {
        case "hi":
            fmt.Println("Hello, how are you")
        case "bye":
            fmt.Println("Goodbye, see you later")
        default:
            fmt.Printf("%s is an unknown input", in)
    }
}