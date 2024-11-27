package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	p := make(map[*NodeL]*NodeL)
	i := l.Head
	for i != l.Tail {
		p[i.Next] = i
		i = i.Next
	}
	blink := &List{}
	for i != l.Head {
		ListPushBack(blink, i.Data)
		i = p[i]
	}
	ListPushBack(blink, i.Data)
	j := l.Head
	k := blink.Head
	for j.Next != nil {
		j.Data = k.Data
		j = j.Next
		k = k.Next
	}
	j.Data = k.Data
}
