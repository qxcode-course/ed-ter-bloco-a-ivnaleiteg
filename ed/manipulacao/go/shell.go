package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var res []int
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			res = append(res, vet[i])
		}
	}

	return res
}

func getCalmWomen(vet []int) []int {
	var res []int
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && abs(vet[i]) < 10 {
			res = append(res, vet[i])
		}
	}

	return res
}

func sortVet(vet []int) []int {
	res := make([]int, len(vet))
	copy(res, vet)

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}

func sortStress(vet []int) []int {
	res := make([]int, len(vet))
	copy(res, vet)

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if abs(res[i]) > abs(res[j]) {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(vet []int) []int {
	res := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		res[i] = vet[len(vet)-1-i]
	}

	return res
}

func unique(vet []int) []int {
	m := make(map[int]bool)
	var res []int

	for i := 0; i < len(vet); i++{
		if !m[vet[i]] {
			m[vet[i]] = true
			res = append(res, vet[i])
		}
	}

	return res
}

func repeated(vet []int) []int {
	count := make(map[int]int)
	var res []int

	for i := 0; i< len(vet); i++ {
		count[vet[i]]++
	}
	
	for i := 0; i < len(vet); i++ {
		if count[vet[i]] > 1 {
			found := false
			for j := 0; j < len(res); j++ {
				if res[j] == vet[i] {
					found = true
				}
			}
			if !found {
				res = append(res, vet[i])
			}
		}
	} 

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

