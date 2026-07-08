package main

import (
	"bufio"
	"fmt"
	"os"
)

func countBattleships(board [][]byte) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}

	count := 0
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 'X' {
				continue
			}

			if i > 0 && board[i-1][j] == 'X' {
				continue
			}

			if j > 0 && board[i][j-1] == 'X' {
				continue
			}

			count++
		}
	}

	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
