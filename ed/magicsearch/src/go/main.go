package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	inicio := 0
	fim := len(slice) - 1

	for inicio <= fim {
		meio := inicio + (fim-inicio)/2

		if slice[meio] == value {
			for meio+1 < len(slice) && slice[meio+1] == value {
				meio++
			}
			return meio
		} else if slice[meio] < value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return inicio
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)

	if len(parts) > 2 {
		for _, elem := range parts[1 : len(parts)-1] {
			value, _ := strconv.Atoi(elem)
			slice = append(slice, value)
		}
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice,value)
	fmt.Println(result)
}