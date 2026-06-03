package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	low := 0
	high := len(slice)

	for low < high {
		mid := low + (high-low)/2

		if slice[mid] == value {
			return true, mid
		} else if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return false, low

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	text := scanner.Text()
	parts := strings.Split(text, " ")
	slice := []int{}

	for _, elem := range parts {
		if elem == "[ " || elem == "] " || elem == "" {
			continue
		}
		value, err := strconv.Atoi(elem)
		if err == nil {
			slice = append(slice, value)
		}
	}
	if !scanner.Scan() {
		return
	}
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
