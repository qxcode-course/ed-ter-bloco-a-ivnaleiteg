package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func naLinha(matriz [][]rune, lin int, num rune) bool {
    for c := 0; c < len(matriz); c++ {
        if matriz[lin][c] == num {
            return true
        }
    }
    return false 
}

func naColuna(matriz [][]rune, col int, num rune) bool {
    for l := 0; l < len(matriz); l++ {
        if matriz[l][col] == num {
            return true
        }
    }
    return false
}

func noQuadrante(matriz [][]rune, lin, col int, num rune) bool {
    dim := len(matriz)

    tamBloco := int(math.Sqrt(float64(dim)))

    inicioLin := (lin / tamBloco) * tamBloco
    inicioCol := (col / tamBloco) * tamBloco

    for l := 0; l < tamBloco; l++ {
        for c := 0; c < tamBloco; c++ {
            if matriz[inicioLin+l][inicioCol+c] == num {
                return true
            }
        }
    }
    return false
}

func resolver(matriz[][]rune, index int) bool {
    nl := len(matriz)

    if index == nl*nl {
        return true
    }

    l := index / nl
    c := index % nl

    if matriz[l][c] != '.' {
        return resolver(matriz, index+1)
    }

    for num := '1'; num < '1'+rune(nl); num++ {
        if !naLinha(matriz, l, num) && !naColuna(matriz, c, num) && !noQuadrante(matriz, l, c, num) {
            matriz[l][c] = num

            if resolver(matriz, index+1) {
                return true
            }
            matriz[l][c] = '.'
        }
    }
    return false
}

func main() {
    var n int
    fmt.Scan(&n)

    scanner := bufio.NewScanner(os.Stdin)
    matriz := make([][]rune, n)

    for i := 0; i < n; i++ {
        scanner.Scan()
        linhaTexto := scanner.Text()
        matriz[i] = []rune(linhaTexto)
    }

    if resolver(matriz, 0) {
        for i := 0; i < n; i++ {
            fmt.Println(string(matriz[i]))
        }
    }
}