package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	p := l
	count := 0
	for count != pos {
		if p == nil {
			return nil
		}
		p = p.Next
		count++
	}
	return p
}
