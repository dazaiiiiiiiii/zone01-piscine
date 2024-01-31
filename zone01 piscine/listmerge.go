package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		return
	}
	p := l1.Head
	for p != nil && p.Next != nil {
		p = p.Next
	}
	p.Next = l2.Head
}
