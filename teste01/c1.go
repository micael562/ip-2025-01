package main

import ("fmt")

func main() {
    var n int
    fmt.Scan(&n)

    soma := 0
    for i := 1; i <= n; i++ {
        fmt.Printf("%d", i) 
        soma += i          
    }
    fmt.Printf("%d", soma)
}
