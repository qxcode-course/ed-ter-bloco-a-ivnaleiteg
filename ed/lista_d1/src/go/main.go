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
	next  *Node 
	prev  *Node 
}

type LList struct {
	root *Node 
	size int   
}


func NewLList() LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return LList{root: root, size: 0}
}


func (ll *LList) Size() int {
	return ll.size
}


func (ll *LList) Clear() {
	if ll.root == nil {
		return
	}
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) insertAfter(prevNode *Node, value int) {
	if prevNode == nil {
		return
	}
	newNode := &Node{Value: value, prev: prevNode, next: prevNode.next}
	prevNode.next.prev = newNode
	prevNode.next = newNode
	ll.size++
}


func (ll *LList) removeNode(node *Node) {
	if node == nil || node == ll.root {
		return 
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--
}

func (ll *LList) PushFront(value int) {
	if ll.root == nil {
		return
	}
	ll.insertAfter(ll.root, value)
}


func (ll *LList) PushBack(value int) {
	if ll.root == nil {
		return
	}
	ll.insertAfter(ll.root.prev, value)
}


func (ll *LList) PopFront() {
	if ll.root == nil || ll.size == 0 {
		return
	}
	ll.removeNode(ll.root.next)
}

func (ll *LList) PopBack() {
	if ll.root == nil || ll.size == 0 {
		return
	}
	ll.removeNode(ll.root.prev)
}


func (ll *LList) String() string {
	if ll.root == nil {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	
	curr := ll.root.next
	for curr != ll.root && curr != nil {
		sb.WriteString(strconv.Itoa(curr.Value))
		if curr.next != ll.root {
			sb.WriteString(", ")
		}
		curr = curr.next
	}
	
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList() 

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, err := strconv.Atoi(v)
				if err != nil {
					continue 
				}
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}