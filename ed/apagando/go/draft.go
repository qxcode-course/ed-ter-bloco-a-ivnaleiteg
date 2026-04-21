package main 

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int 
	fmt.Scan(&m)

	remover := make(map[int]bool)

	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		remover[x] = true 
	}

	resultado := []int{}

	for i := 0; i < n; i++ {
		if !remover[fila[i]] {
			resultado = append(resultado, fila[i])
		}
	}
	
	for i := 0; i < len(resultado); i++ {
		fmt.Print(resultado[i], " ")
	}
	fmt.Println()
}