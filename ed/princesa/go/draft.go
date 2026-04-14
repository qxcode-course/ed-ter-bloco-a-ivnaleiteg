package main

import "fmt"

func main() {
    var n, e int
    fmt.Scan(&n, &e)

    vivos := make([]int, n)
    for i := 0; i < n; i++ {
        vivos[i] = i + 1
    }

    pos := e - 1

        fmt.Print("[ ")
        for i := 0; i < len(vivos); i++ {
            if i == pos {
                fmt.Print(vivos[i], ">")
            } else {
                fmt.Print(vivos[i])
            }
        }
        fmt.Println("] ")

        eliminar := (pos + 1) % len(vivos)
        vivos = append(vivos[:eliminar], vivos[eliminar+1: ]...)

        if eliminar < pos {
            pos --
        }
        pos = pos % len(vivos)
    
    for i := 0; i < len(vivos); i++ {
        if i == pos {
            fmt.Print(vivos[i], ">")
        } else {
            fmt.Print(vivos[i])
        }
        if i < len(vivos)-1 {
            fmt.Print(" ")
        }
    }
        fmt.Println()
}

