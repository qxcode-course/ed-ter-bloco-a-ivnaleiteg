package main

import (
	"fmt"
	"strings"
)

func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for n := l.Front(); n != l.End(); n = n.next {
		sb.WriteString(fmt.Sprint(n.Value))

		if n == sword {
			sb.WriteString(">")
		}
		sb.WriteString(" ")
	}

	res := strings.TrimSuffix(sb.String(), " ")
	return res + " ]"
}

func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil {
			return nil
	}

	next_node := it.next
	if next_node == l.End() {
		next_node = next_node.next
	}
	return next_node
}

func main() {
	var qtd, chosen int

	_, err := fmt.Scanf("%d %d", &qtd, &chosen)
	if err != nil {
		fmt.Scan(&qtd, &chosen)
	}

	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}

	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}

	fmt.Println(ToStr(l, sword))
} 