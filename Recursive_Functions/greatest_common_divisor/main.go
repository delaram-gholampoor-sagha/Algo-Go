package main

import("fmt")

func GCD(m int, n int) int { 
    if m == 0 {
        return n
    }
    if n == 0 {
        return m
    }
    return GCD(n,  m%n) 
}

//Testing code
func main() {
    fmt.Println("GCD is :", GCD(24, 18)) 
}