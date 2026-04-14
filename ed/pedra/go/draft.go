package main

import "fmt"
    
func main() {
    var N int
    fmt.Scan(&N)

    melhor := 1000
    indice := -1

for i := 0; i < N; i++ {
    var A, B int
    fmt.Scan(&A, &B)

    if A < 10 || B < 10{
        continue
    }
    dif := A - B
    if dif < 0 {
        dif = -dif
    }
    if dif < melhor {
        melhor = dif 
        indice = i
    }
}
    if indice == -1{
        fmt.Println("Sem ganhador")
    } else {
        fmt.Println(indice)
    }
}
