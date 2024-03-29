package core

type List struct {
	Car any
	Cdr *List
}

func (l *List) IsEmpty() bool {
	return l.Car == nil && l.Cdr == nil
}

// func (l *List) String() string {
// 	return showList(l)
// }
