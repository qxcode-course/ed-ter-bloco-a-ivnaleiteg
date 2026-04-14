package main

import "fmt"

func main() {
    var N, E, F int
    fmt.Scan(&N, &E, &F)

    vivos := make([]int, N)
    valor := 1 

    for i := 0; i < N; i++ {
        if F == -1 {
            valor = -(i + 1)
            if i%2 == 1 {
                valor = i + 1
            }
    
        } else {
            valor = i + 1
            if i%2 == 1 {
                valor = -(i + 1)
            }
        }
        vivos[i] = valor 
    }
    idx := E % len(vivos)

    for len(vivos) > 1 {
    
        for i, v := range vivos {
            if i == idx {
                if v > 0 {
                    fmt.Printf(">%d ", v)
                } else {
                    fmt.Printf("<%d ", v)
                }
            } else {
                fmt.Printf("%d ", v)
            }
        }
        fmt.Println()

        var alvo int
        
        if vivos[idx] > 0 {
        } else {
            alvo = (idx - 1 + len(vivos)) % len(vivos)
        }

        vivos = append(vivos[:alvo], vivos[alvo+1:]...)
        
        if idx >= len(vivos) {
            idx = 0
        }
    }

    fmt.Println("Vencedor:", vivos[0])
}
