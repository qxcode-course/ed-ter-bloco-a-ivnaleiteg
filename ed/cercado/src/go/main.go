package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)


func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	for i := range board {
		board[i] = bytes.TrimRight(board[i], "\r\n")
	}

	nrows := len(board)
	ncols := len(board[0])

	var dfs func(r, c int)
	
	dfs = func(r, c int) {
		if r < 0 || r >= nrows || c < 0 || c >= ncols || board[r][c] != 'O' {
			return
		}
	
		board[r][c] = 'E'

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	} 

	for c := 0; c < ncols; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[nrows-1][c] == 'O' {
			dfs(nrows-1, c)
		}
	}

	for r := 0; r < nrows; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][ncols-1] == 'O' {
			dfs(r, ncols-1)
		}
	}

	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'E' {
				board[r][c] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}