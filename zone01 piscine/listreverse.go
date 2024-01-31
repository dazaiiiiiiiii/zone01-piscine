package piscine

func ListReverse(l *List) {
	var prev *NodeL
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Tail = l.Head
	l.Head = prev
}
