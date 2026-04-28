package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func (s *Set) Erase(value int) bool {
	idx := s.binarySearch(value)

	if idx >= s.size || s.data[idx] != value {
		return false
	}

	for i := idx; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.size--

	s.data = s.data[:s.size]
	return true
}

func (s *Set) Contains(value int) bool {
	idx := s.binarySearch(value)
	return idx < s.size && s.data[idx] == value
}

func (s *Set) binarySearch(value int) int {
	low := 0
	high := s.size - 1

	for low <= high {
		mid := (low + high) / 2
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, 0, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) find(value int) int {
	for i := 0; i < s.size; i++ {
		if s.data[i] >= value {
			return i
		}
	}
	return s.size
}

func (s *Set) Insert(value int) {
	idx := s.binarySearch(value)

	if idx < s.size && s.data[idx] == value {
		return
	}

	if s.size >= s.capacity {
		if s.capacity == 0 {
			s.capacity = 1
		} else {
			s.capacity *= 2
		}
		
		newData := make([]int, s.size, s.capacity)
		copy(newData, s.data)
		s.data = newData
	}

	s.data = append(s.data, 0)

	for i := s.size; i > idx; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[idx] = value
	s.size++
}

func (s *Set) String() string {
	res := "["
	for i := 0; i < s.size; i++ {
		res += fmt.Sprintf("%v", s.data[i])
		if i < s.size-1 {
			res += ", "
		}
	}
	return res + "]"
}



func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Printf("$%s\n", line)
		
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
		}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "size":
			fmt.Println(v.size)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
