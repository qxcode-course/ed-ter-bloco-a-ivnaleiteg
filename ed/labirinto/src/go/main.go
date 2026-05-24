package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {

	return []Pos{
		{p.l, p.c - 1},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l + 1, p.c},
	}
}

func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func search(grid [][]rune, cur, end Pos) bool {
	if !inside(grid, cur) || grid[cur.l][cur.c] != ' ' {
		return false
	}

	if cur == end {
		grid[cur.l][cur.c] = '.'
		return true
	}

	grid[cur.l][cur.c] = '.'

	for _, next := range getNeig(cur) {
		if search(grid, next, end) {
			return true
		}
	}

	grid[cur.l][cur.c] = ' ' // backtracking
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()

	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)

	grid := make([][]rune, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	var startPos, endPos Pos

	for l := 0; l < nl; l++ {
		for c := 0; c < nc; c++ {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	for _, line := range grid {
		fmt.Println(string(line))
	}
}