package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row int
	col int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	dims := strings.Fields(scanner.Text())
	if len(dims) < 2 {
		return
	}
	rows, _ := strconv.Atoi(dims[0])
	cols, _ := strconv.Atoi(dims[1])

	maze := make([][]rune, rows)
	visited := make([][]bool, rows)

	var start Point
	var end Point

	for r := 0; r < rows; r++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		maze[r] = make([]rune, cols)
		visited[r] = make([]bool, cols)
	
		runes := []rune(line)
		for c := 0; c < cols; c++ {
			if c < len(runes) {
				maze[r][c] = runes[c]
			} else {
				maze[r][c] = ' ' 
			}
		
			if maze[r][c] == 'I' {
				start = Point{row: r, col: c}
			} else if maze[r][c] == 'F' {
				end = Point{row: r, col: c}
			}
		}
	}	
	caminho := NewStack[Point]()
	becos := NewStack[Point]()

	caminho.Push(start)

	dRow := []int{-1, 1, 0, 0}
	dCol := []int{0, 0, -1, 1}
	
	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visited[atual.row][atual.col] = true
			if atual.row == end.row && atual.col == end.col {
				break
		}

		var validos []Point
		for i := 0; i < 4; i++ {
			nRow := atual.row + dRow[i]
			nCol := atual.col + dCol[i]

	
			if nRow >= 0 && nRow < rows && nCol >= 0 && nCol < cols {
				if maze[nRow][nCol] != '#' && !visited[nRow][nCol] {
					validos = append(validos, Point{row: nRow, col: nCol})
				}
			}
		}

		if len(validos) > 0 {
			
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()
		}
	}

	for !caminho.IsEmpty() {
		p := caminho.Pop()
		maze[p.row][p.col] = '.'
	}

	for r := 0; r < rows; r++ {
		fmt.Println(string(maze[r]))
	}
}