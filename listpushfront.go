package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l == nil {
		return
	} else {
		n2 := l.Head
		l.Head = n
		l.Head.Next = n2
		return
	}
}
