package main

import "fmt"

func imprimir (lista []int) {
    saida := fmt.Sprintf ("%v", lista)
    if saida == "[]" {
        fmt.Println ("N")
    } else {
        fmt.Println (saida [1 : len(saida) - 1])
    } 
}

func main () {
    var qtd_album, qtd_possui int
    fmt.Scan (&qtd_album, &qtd_possui)
    figuras := make ([]int, qtd_possui)
    for i := range figuras {
        fmt.Scan (&figuras[i])
    }
    faltando := make([]int, 0)
    unicos := make (map[int]bool)
    repetidos := make ([]int, 0, qtd_possui)

    for _, fig := range figuras {
        if unicos[fig] {
            repetidos = append (repetidos, fig)
        } else {
            unicos[fig] = true
        } 
    }
    for i := 1; i <= qtd_album; i++ {
        if !unicos[i] {
            faltando = append (faltando, i)
        }
    }
    imprimir (repetidos)
    imprimir (faltando)
}