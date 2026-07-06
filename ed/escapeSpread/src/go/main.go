package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 10000000000
const MAX_WAIT = 1000000000

type Point struct {
	r, c int
}

var dirs = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := (len(grid[0]) + 1) / 2

	cleanGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		cleanGrid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cleanGrid[i][j] = int(grid[i][j*2] - '0')
		}
	}

	fireTime := make([][]int, m)
	for i := range fireTime {
		fireTime[i] = make([]int, n)
		for j := range fireTime[i] {
			fireTime[i][j] = INF
		}
	}

	var queue []Point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cleanGrid[i][j] == 1 {
				queue = append(queue, Point{i, j})
				fireTime[i][j] = 0
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nr, nc := curr.r+d.r, curr.c+d.c
			if nr >= 0 && nr < m && nc >= 0 && nc < n && cleanGrid[nr][nc] == 0 {
				if fireTime[nr][nc] == INF {
					fireTime[nr][nc] = fireTime[curr.r][curr.c] + 1
					queue = append(queue, Point{nr, nc})
				}
			}
		}
	}

	canEscape := func(wait int) bool {
		if wait >= fireTime[0][0] {
			return false
		}

		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		playerQueue := []Point{{0, 0}}
		visited[0][0] = true
		time := wait

		for len(playerQueue) > 0 {
			size := len(playerQueue)
			time++
			for i := 0; i < size; i++ {
				curr := playerQueue[0]
				playerQueue = playerQueue[1:]

				if curr.r == m-1 && curr.c == n-1 {
					return true
				}

				for _, d := range dirs {
					nr, nc := curr.r+d.r, curr.c+d.c
					if nr >= 0 && nr < m && nc >= 0 && nc < n && cleanGrid[nr][nc] == 0 && !visited[nr][nc] {
						if nr == m-1 && nc == n-1 {
							if time <= fireTime[nr][nc] {
								visited[nr][nc] = true
								playerQueue = append(playerQueue, Point{nr, nc})
							}
						} else {
							if time < fireTime[nr][nc] {
								visited[nr][nc] = true
								playerQueue = append(playerQueue, Point{nr, nc})
							}
						}
					}
				}
			}
		}
		return false
	}

	if !canEscape(0) {
		return -1
	}
	if canEscape(MAX_WAIT) {
		return MAX_WAIT
	}

	low, high := 0, MAX_WAIT
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		if canEscape(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	var nl, nc int
	fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)

	grid := make([][]byte, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}

	fmt.Println(numIslands(grid))
}