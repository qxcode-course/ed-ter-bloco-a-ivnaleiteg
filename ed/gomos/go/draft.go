package main

import "fmt"

func main() {
	var Q int
	var dir string

	fmt.Scan(&Q, &dir)

	x := make([]int, Q)
	y := make([]int, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	oldX := make([]int, Q)
	oldY := make([]int, Q)

	copy(oldX, x)
	copy(oldY, y)

	switch dir {
	case "L":
		x[0]--
	case "R":
		x[0]++
	case "U":
		y[0]--
	case "D":
		y[0]++
	}

	for i := 1; i < Q; i++ {
		x[i] = oldX[i-1]
		y[i] = oldY[i-1]
	}

	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d\n", x[i],y[i])
	}
}
