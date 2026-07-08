package main

import (
	"fmt"
)

func findStartingStation(gas, cost []int) int {
	totalTank := 0
	currTank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalTank += diff
		currTank += diff

		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}
	return start
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	gas := make([]int, n)
	cost := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&gas[i], &cost[i])
	}

	result := findStartingStation(gas, cost)
	fmt.Println(result)
}