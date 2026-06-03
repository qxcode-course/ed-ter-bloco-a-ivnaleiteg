package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}


func NewMultiSet(capacity int) *MultiSet {
	if capacity == 0 {
		capacity = 2
	}
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}


func (ms *MultiSet) expand() {
	ms.capacity = ms.capacity * 2
	newData := make([]int, ms.capacity)
	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}
	ms.data = newData
}


func (ms *MultiSet) search(value int) (bool, int) {
	low := 0
	high := ms.size

	for low < high {
		mid := low + (high-low)/2
		if ms.data[mid] <= value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	
	if low > 0 && ms.data[low-1] == value {
		return true, low - 1
	}
	return false, low
}


func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	_, index := ms.search(value)
	ms.insert(value, index)
}


func (ms *MultiSet) insert(value int, index int) error {
	if index < 0 || index > ms.size {
		return errors.New("index out of bounds")
	}
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
	return nil
}


func (ms *MultiSet) Erase(value int) error {
	found, index := ms.search(value)
	if !found {
		return errors.New("value not found")
	}
	return ms.erase(index)
}


func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= ms.size {
		return errors.New("index out of bounds")
	}
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}


func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}


func (ms *MultiSet) Count(value int) int {
	found, index := ms.search(value)
	if !found {
		return 0
	}
	count := 0
	for i := index; i >= 0; i-- {
		if ms.data[i] == value {
			count++
		} else {
			break
		}
	}
	return count
}


func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	uniqueCount := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}


func (ms *MultiSet) Clear() {
	ms.size = 0
}


func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < ms.size; i++ {
		sb.WriteString(strconv.Itoa(ms.data[i]))
		if i < ms.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		line = scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		
		fmt.Println("$" + line)
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}