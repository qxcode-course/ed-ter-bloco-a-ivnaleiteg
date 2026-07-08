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
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	curr := l.root.next
	first := true
	for curr != l.root {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(curr.Value))
		first = false
		curr = curr.next
	}
	sb.WriteString("]")
	return sb.String()
}

func equals(lla *LList, llb *LList) bool {
	currA := lla.root.next
	currB := llb.root.next

	for currA != lla.root && currB != llb.root {
		if currA.Value != currB.Value {
			return false
		}
		currA = currA.next
		currB = currB.next
	}

	return currA == lla.root && currB == llb.root
}

func addsorted(ll *LList, value int) {
	curr := ll.root.next
	for curr != ll.root && curr.Value < value {
		curr = curr.next
	}
	ll.insertBefore(curr, value)
}

func reverse(ll *LList) {
	if ll.size <= 1 {
		return
	}
	curr := ll.root
	for {
		curr.next, curr.prev = curr.prev, curr.next
		curr = curr.prev
		if curr == ll.root {
			break
		}
	}
}

func merge(lla *LList, llb *LList) *LList {
	result := NewLList()
	currA := lla.root.next
	currB := llb.root.next

	for currA != lla.root && currB != llb.root {
		if currA.Value <= currB.Value {
			result.PushBack(currA.Value)
			currA = currA.next
		} else {
			result.PushBack(currB.Value)
			currB = currB.next
		}
	}

	for currA != lla.root {
		result.PushBack(currA.Value)
		currA = currA.next
	}

	for currB != llb.root {
		result.PushBack(currB.Value)
		currB = currB.next
	}

	return result
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(strings.TrimSpace(p))
		ll.PushBack(value)
	}
	return ll
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}