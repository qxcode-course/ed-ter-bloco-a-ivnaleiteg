package main

import "fmt"

type Ponto struct {
    linha  int
    coluna int
}

func orangesRotting(grid [][]int) int {
    linhas := len(grid)
    if linhas == 0 {
        return 0
    }
    colunas := len(grid[0])

    var fila []Ponto
    laranjasFrescas := 0

    for l := 0; l < linhas; l++ {
        for c := 0; c < colunas; c++ {
            if grid[l][c] == 2 {
                fila = append(fila, Ponto{linha: l, coluna: c})
            } else if grid[l][c] == 1 {
                laranjasFrescas++
            }
        }
    }

    if laranjasFrescas == 0 {
        return 0
    }

    minutos := 0
    direcoes := [][]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    for len(fila) > 0 && laranjasFrescas > 0 {
        minutos++
        tamanhoRodada := len(fila)

        for i := 0; i < tamanhoRodada; i++ {
            atual := fila[0]
            fila = fila[1:]

            for _, dir := range direcoes {
                novaL := atual.linha + dir[0]
                novaC := atual.coluna + dir[1]

                if novaL >= 0 && novaL < linhas && novaC >= 0 && novaC < colunas {
                    if grid[novaL][novaC] == 1 {
                        grid[novaL][novaC] = 2
                        laranjasFrescas--
                        fila = append(fila, Ponto{linha: novaL, coluna: novaC})
                    }
                }
            }
        }
    }

    if laranjasFrescas > 0 {
        return -1
    }
    return minutos
}


func main() {
    var m, n int

    for {
        _, err := fmt.Scan(&m, &n)
        if err != nil {
            break
        }

        grid := make([][]int, m)
        for i := 0; i < m; i++ {
            grid[i] = make([]int, n)
            for j := 0; j < n; j++ {
                fmt.Scan(&grid[i][j])
            }
        }
        fmt.Println(orangesRotting(grid))
    }
}
    
  