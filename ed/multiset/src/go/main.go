package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func (ms *MultiSet) Clear() {
	ms.data = ms.data[:0]
	ms.size = 0
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Count(value int) int {
	found, idx := ms.search(value)
	if !found {
		return 0
	}
	count := 0
	for i := idx; i < ms.size && ms.data[i] == value; i++ {
		count++
	}
	for i := idx - 1; i >= 0 && ms.data[i] == value; i-- {
		count++
	}
	return count
}

func (ms *MultiSet) search(value int) (bool, int) {
	low := 0
	high := ms.size - 1
	for low <= high {
		mid := (low + high) / 2
		if ms.data[mid] == value {
			return true, mid
		}
		if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false, low
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	_, idx := ms.search(value)
	ms.data = append(ms.data, 0)
	for i := ms.size; i > idx; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[idx] = value
	ms.size++
}

func (ms *MultiSet) Erase(value int) bool {
	found, idx := ms.search(value)
	if !found {
		return false
	}
	ms.erase(idx)
	return true
}

func (ms *MultiSet) erase(index int) {
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	ms.data = ms.data[:ms.size]
}

func (ms *MultiSet) Contains(value int) bool {
    found, _ := ms.search(value)
    return found
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, 0, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}
	newData := make([]int, ms.size, ms.capacity)
	copy(newData, ms.data)
	ms.data = newData
}

func (ms *MultiSet) String() string {
	res := "["
	for i := 0; i < ms.size; i++ {
		res += strconv.Itoa(ms.data[i])
		if i < ms.size-1 {
			res += ", "
		}
	}
	return res + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			continue
		}
		fmt.Printf("$%s\n", line)
		args := strings.Fields(line)
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
			if !ms.Erase(value) {
				fmt.Println("value not found")
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
