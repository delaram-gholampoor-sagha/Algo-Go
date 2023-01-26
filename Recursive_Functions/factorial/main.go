package main

import("fmt")

func Factorial(i int) int {
    // Termination Condition
    if i <= 1 {
        return 1 
    }
    // Body, Recursive Expansion
    return i * Factorial(i-1)
}
//Testing code
func main() {
    fmt.Println("factorial 6 is :: ", Factorial(-6))
}