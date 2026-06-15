package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

type Iterator struct {
	data  []int
	index int
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	if !i.HasNext() {
		panic(fmt.Errorf("no more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

type ReverseIterator struct {
	data  []int
	index int
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{data: l.data, index: len(l.data)}
}

func (i *ReverseIterator) HasNext() bool {
	return i.index > 0
}

func (i *ReverseIterator) Next() int {
	if !i.HasNext() {
		panic(fmt.Errorf("No more elements"))
	}
	i.index -= 1
	return i.data[i.index]
}

type CyclicIterator struct {
	data  []int
	index int
}

func (l *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{data: l.data, index: -1}
}

func (i *CyclicIterator) HasNext() bool {
	return len(i.data) > 0
}

func (i *CyclicIterator) Next() int {
	if !i.HasNext() {
		panic(fmt.Errorf("no elements to iterate"))
	}
	i.index += 1
	if i.index >= len(i.data) {
		i.index = 0
	}
	return i.data[i.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			slice := make([]int, len(args)-1)
			for i, value := range args[1:] {
				slice[i], _ = strconv.Atoi(value)
			}
			mylist = NewMyList(slice)
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			if len(args) < 2 {
				continue
			}
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			if it.HasNext() {
				for range qtd {
					fmt.Printf("%v ", it.Next())
				}
			}
			fmt.Println("]")
		}
	}
}


