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
	root  *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{
		root: root,
		size: 0,
	}
}

func (ll *LList) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.root.prev == ll.root {
		return nil
	}
	return ll.root.prev
}

func (ll *LList) Search(value int) *Node {
	node := ll.root.next
	for node != ll.root {
		if node.Value == value {
			return node
		}
		node = node.next
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	newNode := &Node{Value: value, root: ll.root}
	prevNode := node.prev

	newNode.next = node
	newNode.prev = prevNode

	prevNode.next = newNode
	node.prev = newNode

	ll.size++
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}

	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode

	ll.size--

	if nextNode == ll.root {
		return nil
	}
	return nextNode
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{Value: value, root: ll.root}
	first := ll.root.next

	newNode.next = first
	newNode.prev = ll.root

	ll.root.next = newNode
	first.prev = newNode

	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{Value: value, root: ll.root}
	last := ll.root.prev

	newNode.next = ll.root
	newNode.prev = last

	last.next = newNode
	ll.root.prev = newNode

	ll.size++
}

func (ll *LList) String() string {
	var values []string
	node := ll.root.next

	for node != ll.root {
		values = append(values, strconv.Itoa(node.Value))
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
			 //ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			 ll.Clear()
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.Value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Insert(node, newvalue)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "remove":
			 oldvalue, _ := strconv.Atoi(args[1])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Remove(node)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
