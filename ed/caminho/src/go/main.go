package main

import (
	"bufio"
	"fmt"
	"os"
	
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{l: p.l - 1, c: p.c},
		{l: p.l + 1, c: p.c}, 
		{l: p.l, c: p.c - 1},
		{l: p.l, c: p.c + 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := NewQueue[Pos]()
	queue.Enqueue(startPos)

	visitados := make(map[Pos]bool)
	visitados[startPos] = true

	caminho := make(map[Pos]Pos)

	destinoEncontrado := false

	for !queue.IsEmpty() {
		atual, _ := queue.Dequeue()

		if atual == endPos {
			destinoEncontrado = true
			break
		}

		for _, vizinho := range atual.getNeig() {
			if inside(grid, vizinho) && !visitados[vizinho] && match(grid, vizinho, ' ') {
				visitados[vizinho] = true
				queue.Enqueue(vizinho)
				caminho[vizinho] = atual
			}
		}
	}

	if destinoEncontrado {
		voltar(caminho, endPos, startPos, grid)
	}
}

func voltar(caminho map[Pos]Pos, atual Pos, startPos Pos, grid [][]rune) {
	grid[atual.l][atual.c] = '.'

	for atual != startPos {
		atual = caminho[atual]
		grid[atual.l][atual.c] = '.'
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
