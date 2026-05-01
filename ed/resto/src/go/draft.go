package main

import "fmt"

func divResto(n int) {
    if n == 0 {
        return
    }

    q := n / 2
    r := n % 2

    divResto(q)

    fmt.Printf("%d %d\n", q, r)
}
func main() {
    var n int 
    fmt.Scan(&n)

    divResto(n)
}
