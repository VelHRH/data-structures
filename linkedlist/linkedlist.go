package linkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) Prepend(v int) {
	newNode := &node{
		value: v,
		next:  l.head,
	}
	l.head = newNode
	l.length++
}

func (l LinkedList) Print() {
	currNode := l.head
	for l.length > 0 {
		fmt.Printf("%d ", currNode.value)
		currNode = currNode.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) DeleteByValue(v int) {
	currNode := l.head
	if currNode.value == v {
		l.head = currNode.next
		l.length--
		return
	}
	for currNode.next.value != v {
		currNode = currNode.next
		if currNode.next == nil {
			fmt.Printf("There's no value %d \n", v)
			return
		}
	}
	currNode.next = currNode.next.next
	l.length--
}
