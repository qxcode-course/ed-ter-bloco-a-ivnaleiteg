package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	if capacity < 1 {
		capacity = 1
	}
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, 0, newCapacity)
	newData = append(newData, s.data[:s.size]...)
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	left, right := 0, s.size-1

	for left <= right {
		mid := left + (right-left)/2
		if mid >= s.size {
			break
		}
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func (s *Set) findInsertionIndex(value int) int {
	left, right := 0, s.size-1

	for left <= right {
		mid := left + (right-left)/2
		if mid >= s.size {
			break
		}
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func (s *Set) insert(value int, index int) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("index out of bounds")
	}

	if s.size >= s.capacity {
		s.reserve(s.capacity * 2)
	}

	s.data = append(s.data, 0)

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[index] = value
	s.size++
	return nil
}

func (s *Set) erase(index int) error {
	if index < 0 || index >= s.size {
		return fmt.Errorf("index out of bounds")
	}

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	
	s.size--
	s.data = s.data[:s.size]
	return nil
}

func (s *Set) Insert(value int) {
	idx := s.findInsertionIndex(value)
	if idx < s.size && s.data[idx] == value {
		return
	}
	_ = s.insert(value, idx)
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) != -1
}

func (s *Set) Erase(value int) bool {
	idx := s.binarySearch(value)
	if idx == -1 {
		return false
	}
	_ = s.erase(idx)
	return true
}

func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.size; i++ {
		sb.WriteString(strconv.Itoa(s.data[i]))
		if i < s.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var set *Set = nil

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		
		cmd = strings.TrimPrefix(parts[0], "$")

		switch cmd {
		case "end":
			return
            
		case "init":
			capacity, _ := strconv.Atoi(parts[1])
			set = NewSet(capacity)

		case "insert":
			if set != nil {
				for _, part := range parts[1:] {
					value, _ := strconv.Atoi(part)
					set.Insert(value)
				}
			}
            
		case "show":
			if set != nil {
				fmt.Println(set.String())
			} else {
				fmt.Println("[]")
			}
            
		case "erase":
			if set != nil {
				value, _ := strconv.Atoi(parts[1])
				if !set.Erase(value) {
					fmt.Println("value not found")
				}
			}
            
		case "contains":
			if set != nil {
				value, _ := strconv.Atoi(parts[1])
				if set.Contains(value) {
					fmt.Println("true")
				} else {
					fmt.Println("false")
				}
			}
            
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}