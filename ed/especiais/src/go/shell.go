package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	freq := make(map[int]int)
	for _, v := range vet {
		if v < 0 {
			v = -v
		}
		freq[v]++
	}

	var res []Pair
	for k, v := range freq {
		res = append(res, Pair{k, v})
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i].One > res[j].One {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	var res []Pair
	atual := vet[0]
	count := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			count++
		} else {
			res = append(res, Pair{atual, count})
			atual = vet[i]
			count = 1
		}
	}

	res = append(res, Pair{atual, count})
	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if (i > 0 && vet[i-1] < 0) || (i < len(vet)-1 && vet[i+1] < 0) {
				res[i] = 1
			}
		}
	}

	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			left := i > 0 && vet[i-1] < 0 
			right := i < len(vet)-1 && vet[i+1] < 0

			if !left && !right {
				res[i] = 1
			}
		}
	}
	return res
}

func couple(vet []int) int {
	homens := make(map[int]int)
	mulheres := make(map[int]int)

	for _, v := range vet {
		if v > 0 {
			homens[v]++
		} else {
			mulheres[-v]++
		}
	}

	total := 0
	for stress, h := range homens {
		m := mulheres[stress]
		if h < m {
			total += h
		} else {
			total += m
		}
	}
	return total
}
func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}

	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}

	return true
}

func subseq(vet []int, seq []int) int {
	for i := 0; i < len(vet); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	remover := make(map[int]bool)
	for _, p := range posList {
		remover[p] = true
	}

	var res []int
	for i := 0; i < len(vet); i++ {
		if !remover[i] {
			res = append(res, vet[i])
		}
	}

	return res
}

func clear(vet []int, value int) []int {
	var res []int 
	for _, v := range vet {
		if v < 0 {
			if -v != value {
				res = append(res, v)
			}
		} else {
			if v != value {
				res = append(res, v)
			}
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlicePair(occurr(str2vet(args[1])))
		case "teams":
			printSlicePair(teams(str2vet(args[1])))
		case "mnext":
			printSliceInt(mnext(str2vet(args[1])))
		case "alone":
			printSliceInt(alone(str2vet(args[1])))
		case "erase":
			printSliceInt(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSliceInt(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSliceInt(vet []int) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func printSlicePair(vet []Pair) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}