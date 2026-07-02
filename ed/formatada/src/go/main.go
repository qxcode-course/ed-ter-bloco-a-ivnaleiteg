package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func MyShow(node *Node, nivel int) {
	if node == nil {
		return
	}

	prefixo := strings.Repeat("....", nivel)
	if node.Left == nil && node.Right == nil {
		fmt.Printf("%s%d\n", prefixo, node.Value)
		return
	}

	if node.Left != nil {
		MyShow(node.Left, nivel+1)
	} else {
		fmt.Printf("%s#\n", strings.Repeat("....", nivel+1))
	}
	fmt.Printf("%s%d\n", prefixo, node.Value)

	if node.Right != nil {
		MyShow(node.Right, nivel+1)
	} else {
		fmt.Printf("%s#\n", strings.Repeat("....", nivel+1))
	}
}

func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
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
		BShow(node.Right, history+"r")
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
	BShow(root, "") 
	MyShow(root, 0)
}
