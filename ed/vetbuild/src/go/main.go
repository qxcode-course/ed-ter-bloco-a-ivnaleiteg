package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}

	newData := make([]int, newCapacity)

	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}

	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		if v.capacity == 0 {
			v.Reserve(1)
		} else {
			v.Reserve(v.capacity * 2)
		}
	}

	v.data[v.size] = value
	v.size++
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) PopBack() (int, error) {
	if v.size == 0 {
		return 0, errors.New("vector is empty")
	}

	value := v.data[v.size-1]
	v.size--
	return value, nil
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) String() string {
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

func (v *Vector) Get(index int) int {
	return v.data[index]
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, errors.New("index out of range")
	}

	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return errors.New("index out of range")
	}

	v.data[index] = value
	return nil
}

func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
		return errors.New("index out of range")
	}

	if v.size == v.capacity {
		if v.capacity == 0 {
			v.Reserve(1)
		} else {
			v.Reserve(v.capacity * 2)
		}
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++

	return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return errors.New("index out of range")
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--

	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}

	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) Slice(start int, end int) *Vector {
	if start < 0 {
		start = v.size + start
	}
	if end < 0 {
		end = v.size + end
	}

	if start < 0 {
		start = 0
	}
	if end > v.size {
		end = v.size
	}
	if start > end {
		start = end
	}

	return &Vector{
		data:     v.data[start:end],
		size:     end - start,
		capacity: end - start,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)

	for {
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			fmt.Println(line)
			continue
		}

		if !strings.HasPrefix(line, "$") {
			fmt.Println("$" + line)
		} else {
			fmt.Println(line)
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd = strings.TrimPrefix(parts[0], "$")

		switch cmd {
		case "end":
			return

		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)

		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}

		case "show":
			fmt.Println(v)

		case "status":
			fmt.Println(v.Status())

		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}

		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}

		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)

		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

		case "clear":
			v.Clear()

		case "capacity":
			fmt.Println(v.Capacity())

		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}

		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)

		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}