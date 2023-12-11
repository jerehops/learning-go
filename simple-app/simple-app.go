package main
 	
import (
    "fmt"
    "math/rand"
)
func main() {
    quote := []string{
        "An apple a day keeps the doctor away",
        "Knowing is not enough, we must apply. Willing is not enough, we must do",
        "Life is full of mystery",
        "Yesterday you said today. It is already tomorrow",
        
    }
    var input int
    var value bool
    for {
        fmt.Println("What would you like to do?")
        fmt.Println("1) Get a random quote")
        fmt.Println("2) Do simple math")
        fmt.Println("3) Quit")
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&input)
        switch input {
            case 1 :
                min := 0
                max := 3
                fmt.Println("Your random quote of the day is: ")
                fmt.Printf("%s\n\n", quote[rand.Intn(max - min) + min])
            case 2 :
                var num1, num2 int
                var arithmetic string
                fmt.Print("Your first number is: ")
                fmt.Scan(&num1)
                fmt.Print("Your second number is: ")
                fmt.Scan(&num2)
                for {
                    fmt.Print("What do you want to do: ")
                    fmt.Scan(&arithmetic)
                    switch arithmetic {
                        case "+" :
                            fmt.Printf("%d %s %d = %d\r\n", num1, arithmetic, num2, (num1+num2))
                            value = true
                        case "-" :
                            fmt.Printf("%d %s %d = %d\r\n", num1, arithmetic, num2, (num1-num2))
                            value = true
                        case "*" :
                            fmt.Printf("%d %s %d = %d\r\n", num1, arithmetic, num2, (num1*num2))
                            value = true
                        case "/" :
                            fmt.Printf("%d %s %d = %d\r\n", num1, arithmetic, num2, (num1/num2))
                            value = true                            
                        default :
                            fmt.Printf("Invalid option, only + - / * allowed\r\n")
                            value = false
                    }
                    if value == true {
                        break
                    }
                }
            case 3 :
                fmt.Println("Thank you come again.")
                value = true
            default :
                fmt.Println("Invalid option")
        }
        if value == true {
            break
        }
    }
}