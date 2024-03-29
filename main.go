package main

import (
	"corelisp/core"
	"fmt"
)

func main() {
	l := core.Parce("(CDR (QUOTE ((* 5 7) 3)))")
	// fmt.Println(l)
	print(l)
}

func print(l *core.List) {
	for !l.IsEmpty() {
		ll, ok := l.Car.(*core.List)
		if ok {
			print(ll)
		} else {
			fmt.Println(l.Car)
		}
		l = l.Cdr
	}
}
