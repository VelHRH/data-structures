type ListNode struct {
	Next  *ListNode
	value int
}

func (l *ListNode) Prepend(v int) {
	newNode := ListNode{value: v, Next: l.Next}
	l.Next = &newNode
}

func (l *ListNode) Print() {

	for l.Next != nil {
		fmt.Printf("%d ", l.value)
		l = l.Next
	}
	fmt.Println(l.value)
}
