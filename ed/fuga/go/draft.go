package main

import "fmt"

func main() {
	var H, P, F, D int

	fmt.Scan(&H, &P, &F, &D)

	for {
		F = F + D

		if F < 0 {
			F = 15
		}
		if F > 15 {
			F = 0
		}
		if F == P {
			fmt.Printf("N")
			return
		}
		if F == H {
			fmt.Printf("S")
			return
		}
	}
} 