package main

import "fmt"

func ehValido(sequencia []rune, index int, digito rune, L int) bool {
    for i := 1; i <= L; i++ {
        if index-i >= 0 && sequencia[index-i] == digito {
            return false
        }
        if index+i < len(sequencia) && sequencia[index+i] == digito {
            return false 
        }
    }
    return true
} 

func resolverBacktracking(sequencia []rune, L int) ([]rune, bool) {
    index := -1 
    for i, r := range sequencia {
        if r == '.' {
            index = i
            break
        }
    }
    if index == -1 {
        return sequencia, true
    }
    for d := 0; d <= L; d++ {
        digitoRune := rune('0' + d)

        if ehValido(sequencia, index, digitoRune, L) {
            sequencia[index] = digitoRune 

            if resultado, resolvido := resolverBacktracking(sequencia, L); resolvido {
                return resultado, true 
            }
            sequencia[index] = '.'
        }
    }
    return nil, false
}

func main() {
    var sequenciaStr string
    var L int

    if _, err := fmt.Scan(&sequenciaStr); err != nil {
        return
    }
    if _, err := fmt.Scan(&L); err != nil {
        return
    }

    sequencia := []rune(sequenciaStr)

    if solucao, resolvido := resolverBacktracking(sequencia, L); resolvido {
            fmt.Println(string(solucao))
    }
}
