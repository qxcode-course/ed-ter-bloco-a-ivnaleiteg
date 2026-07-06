package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	corOriginal := image[sr][sc]

	if corOriginal == color {
		return image
	}

	linhas := len(image)
	colunas := len(image[0])

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || r >= linhas || c < 0 || c >= colunas || image[r][c] != corOriginal {
			return 
		}

		image[r][c] = color

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	
	dfs(sr, sc)

	return image
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
