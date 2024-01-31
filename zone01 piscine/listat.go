package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil && count < pos {
		l = l.Next
		count++
	}
	return l
}
