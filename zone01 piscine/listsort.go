package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	curr := l
	for ; curr != nil; curr = curr.Next {
		for next := curr.Next; next != nil; next = next.Next {
			if curr.Data > next.Data {
				curr.Data, next.Data = next.Data, curr.Data
			}
		}
	}
	return l
}
