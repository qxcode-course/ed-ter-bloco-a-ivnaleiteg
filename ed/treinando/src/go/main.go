package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(int) string
	rec = func(i int) string {
		if i == len(vet)-1 {
			return fmt.Sprintf("%d", vet[i])
		}
		return fmt.Sprintf("%d, ", vet[i]) + rec(i+1)
	}

	return "[" + rec(0) + "]"
}
func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(int) string
	rec = func(i int) string {
		if i == 0 {
			return fmt.Sprintf("%d", vet[i])
		}
		return fmt.Sprintf("%d, ", vet[i]) + rec(i-1)
	}

	return "[" + rec(len(vet)-1) + "]"
}

func reverse(vet []int) {
	var rec func(int, int)
	rec = func(i int, j int) {
		if i >= j {
			return
		}
		vet[i], vet[j] = vet[j], vet[i]
		rec(i+1, j-1)
	}
	rec(0, len(vet)-1)
}

func sum(vet []int) int {
	var rec func(int) int
	rec = func(i int) int {
		if i == len(vet) {
			return 0
		}
		return vet[i] + rec(i+1)
	}
	return rec(0)
}

func mult(vet []int) int {
	var rec func(int) int
	rec = func(i int) int {
		if i == len(vet) {
			return 1
		}
		return vet[i] * rec(i+1)
	}
	return rec(0)
}

func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(int) (int, int)
	rec = func(i int) (int, int) {
		if i == len(vet)-1 {
			return i, vet[i]
		}

		idx, val := rec(i + 1)

		if vet[i] < val {
			return i, vet[i]
		}
		return idx, val
	}

	idx, _ := rec(0)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}