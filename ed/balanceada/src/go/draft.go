package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isBalanced(s string) bool {
    var stack []rune

    for _, char := range s {
        if char == '(' || char == '[' {
            stack = append(stack, char)
        } else if char == ')' || char == ']' {
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]

            if (char == ')' && top == '(') || (char == ']' && top == '[') {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        input := strings.TrimSpace(scanner.Text())

        if isBalanced(input) {
            fmt.Println("balanceado")
        } else {
            fmt.Println("nao balanceado")
        }
    }
}
