package piscine

func ListForEach(l *List, f func(*NodeL)) {
	p := l.Head
	for p != nil {
		f(p)
		p = p.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}
