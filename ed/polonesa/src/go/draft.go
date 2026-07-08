package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

func isOperator(token string) bool {
	return precedence(token) > 0
}

func toRPN(tokens []string) string {
	var output []string
	var stack []string

	for _, token := range tokens {
		if isOperator(token) {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if (token != "^" && precedence(top) >= precedence(token)) || (token == "^" && precedence(top) > precedence(token)) {
					output = append(output, top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, token)
		} else {
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return strings.Join(output, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		fmt.Println(toRPN(tokens))
	}
}