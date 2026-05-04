package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{
		root: root,
		size: 0,
	}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{Value: value}
	first := ll.root.next

	newNode.next = first
	newNode.prev = ll.root

	ll.root.next = newNode
	first.prev = newNode

	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{Value: value}
	last := ll.root.prev

	newNode.next = ll.root
	newNode.prev = last

	last.next = newNode
	ll.root.prev = newNode

	ll.size++
}

func (ll *LList) PopFront() {
	if ll.size == 0 {
		return
	}

	first := ll.root.next
	second := first.next

	ll.root.next = second
	second.prev = ll.root

	ll.size--
}

func (ll *LList) PopBack() {
	if ll.size == 0 {
		return
	}

	last := ll.root.prev
	penultimate := last.prev

	penultimate.next = ll.root
	ll.root.prev = penultimate

	ll.size--
}

func (ll *LList) String() string {
	var values []string
	node := ll.root.next

	for node != ll.root {
		values = append(values, fmt.Sprintf("%d", node.Value))
		node = node.next
	}

	return "[" + strings.Join(values, ", ") + "]"
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
			num, _ := strconv.Atoi(v)
			ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
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
