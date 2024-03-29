package core

import (
	"strings"
)

func Parce(s string) *List {
	l, _ := parce(s)
	return l
}

func parce(s string) (*List, int) {
	r := new(List)
	t := r
	left_p := 1
	right_p := 0
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] == '(' {
			left_p += 1
			car, sd := parce(s[i+1:])
			i += (sd + 1)
			t.Car = car
		} else if s[i] == ')' {
			right_p += 1

		} else {
			t.Car = parceAtom(s[i:])
		}

		t.Cdr = new(List)
		t = t.Cdr
	}

	return r, i
}

func parceAtom(s string) string {
	builder := new(strings.Builder)
	for _, c := range s {
		if c == ' ' || c == ')' {
			break
		}
		if c == ')' {
			continue
		}
		builder.WriteRune(c)
	}
	return builder.String()
}
