package piscine

func ListSize(l *List) int {
	if l.Head == l.Tail {
		return 0
	}
	count := 0
	p := l.Head
	for p != l.Tail {
		count++
		p = p.Next
	}
	count++
	return count
}
