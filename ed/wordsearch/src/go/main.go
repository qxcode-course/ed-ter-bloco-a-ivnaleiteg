package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}

	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == word[0] {
				if dfs(grid, r, c, word, 0) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(grid [][]byte, r int, c int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}

	temp := grid[r][c]
	grid[r][c] = '*'

	dRow := []int{-1, 1, 0, 0}
	dCol := []int{0, 0, -1, 1}

	for i := 0; i < 4; i++ {
		if dfs(grid, r+dRow[i], c+dCol[i], word, index+1) {
			return true
		}
	}

	grid[r][c] = temp

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
