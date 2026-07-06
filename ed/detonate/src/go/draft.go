package main

import "fmt"

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	grafo := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			x1, y1, r1 := int64(bombs[i][0]), int64(bombs[i][1]), int64(bombs[i][2])
			x2, y2 := int64(bombs[j][0]), int64(bombs[j][1])

			dx := x1 - x2
			dy := y1 - y2
			distanciaQuadrada := (dx * dx) + (dy * dy)

			raioQuadrado := r1 * r1
			
			if distanciaQuadrada <= raioQuadrado {
				grafo[i] = append(grafo[i], j)
			}
		}
	}

	maxBombas := 0

	for i := 0; i < n; i++ {
		detonadas := bfs(i, n, grafo)
		if detonadas > maxBombas {
			maxBombas = detonadas
		}
	}

	return maxBombas
}

func bfs(inicio int, n int, grafo [][]int) int {
	visitado := make([]bool, n)
	var fila []int

	fila = append(fila, inicio)
	visitado[inicio] = true
	contagem := 1

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		for _, vizinho := range grafo[atual] {
			if !visitado[vizinho] {
				visitado[vizinho] = true
				contagem++
				fila = append(fila, vizinho)
			}
		}
	}

	return contagem
}

func main() {
	var n, col int

	for {
		_, err := fmt.Scan(&n, &col)
		if err != nil {
			break
		}

		bombs := make([][]int, n)
		for i := 0; i < n; i++ {
			bombs[i] = make([]int, col)
			for j := 0; j < col; j++ {
				fmt.Scan(&bombs[i][j])
			}
		}

		fmt.Println(maximumDetonation(bombs))
	}
}