package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, r int, c int) {
	rows := len(grid)
	cols := len(grid[0])

	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func numIslands(grid[][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++

				dfs(grid, r, c)
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}