package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	stack := NewStack[Pos]()
	stack.Push(Pos{l: l, c: c})

	nl := len(grid)
	nc := len(grid[0])

	for !stack.IsEmpty() {
		pos := stack.Pop()

		if pos.l < 0 || pos.l >= nl || pos.c < 0 || pos.c >= nc {
			continue
		}

		if grid[pos.l][pos.c] != '#' {
			continue
		}

		grid[pos.l][pos.c] = 'o'

		stack.Push(Pos{l: pos.l - 1, c: pos.c})
		stack.Push(Pos{l: pos.l + 1, c: pos.c})
		stack.Push(Pos{l: pos.l, c: pos.c - 1})
		stack.Push(Pos{l: pos.l, c: pos.c + 1})
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}