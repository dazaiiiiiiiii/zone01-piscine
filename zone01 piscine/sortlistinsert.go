package piscine

func listPushBack(l *NodeI, data int) *NodeI {
	n := NodeI{
		Data: data,
	}
	if l == nil {
		return &n
	}
	move := l
	for move.Next != nil {
		move = move.Next
	}
	move.Next = &n
	return l
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	link := listPushBack(l, data_ref)
	ListSort(link)
	return link
}
