package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func rec_sum(node *Node) int {
	if node == nil {
		return 0

	}
			return node.Value + rec_sum(node.Left) + rec_sum(node.Right)
}

func rec_min(node *Node) int {
	if node == nil {
		return math.MaxInt
	}

	minVal := node.Value

	if leftMin := rec_min(node.Left); leftMin < minVal {
		minVal = leftMin
	}
	if rightMin := rec_min(node.Right); rightMin < minVal {
		minVal = rightMin
	}

	return minVal
}

 
func MyShow(node *Node, nivel int) {
	_, _ = node, nivel
}

func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, heranca+"r")
	}
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	root := create(&parts)
	fmt.Println("Arvore:")
	BShow(root, "")
	fmt.Printf("Soma: %d, Minimo: %d\n", rec_sum(root), rec_min(root))
}
