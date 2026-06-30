package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)


func ToStr(l *list.List, sword *list.Element) string {
	var elements []string
	for it := l.Front(); it != nil; it = it.Next() {
		val := it.Value.(int)
		if it == sword {
			if val > 0 {
				elements = append(elements, strconv.Itoa(val)+">")
			} else {
				elements = append(elements, "<"+strconv.Itoa(val))
			}
		} else {
			elements = append(elements, strconv.Itoa(val))
		}
	}
	return "[ " + strings.Join(elements, " ") + " ]"
}
func Next(l *list.List, it *list.Element) *list.Element {
	if it == nil || it.Next() == nil {
		return l.Front()
	}
	return it.Next()
}


func Prev(l *list.List, it *list.Element) *list.Element {
	if it == nil || it.Prev() == nil {
		return l.Back()
	}
	return it.Prev()
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}