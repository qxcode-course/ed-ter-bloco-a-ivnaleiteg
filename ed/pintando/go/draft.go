package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Os lados devem ser números positivos")
		return
	}

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Os valores não formam um triângulo válido")
		return
	}

	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	fmt.Printf("%.2f\n", area)
}