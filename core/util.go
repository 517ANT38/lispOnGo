package core

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"strings"
// )

// // func BinOp(op spechForm, ina float64, inb float64) float64 {
// // 	switch op {
// // 	case AR_ADD:
// // 		return ina + inb
// // 	case AR_SUB:
// // 		return ina - inb
// // 	case AR_MUL:
// // 		return ina * inb
// // 	case AR_DIV:
// // 		return ina / inb
// // 	case AR_MOD:
// // 		return math.Mod(ina, inb)
// // 	default:
// // 		return ina
// // 	}
// // }

// // func CompOp(op spechForm, ina float64, inb float64) bool {
// // 	switch op {
// // 	case AR_GR:
// // 		return ina > inb
// // 	case AR_GR_EQ:
// // 		return ina >= inb
// // 	case AR_LOW:
// // 		return ina < inb
// // 	case AR_LOW_EQ:
// // 		return ina <= inb
// // 	case AR_EQ:
// // 		return Equals(ina, inb)
// // 	case AR_NOT_EQ:
// // 		return ina != inb
// // 	default:
// // 		return true
// // 	}
// // }

// // func Equals(numA, numB float64) bool {
// // 	delta := math.Abs(numA - numB)
// // 	return delta < 1e-10
// // }
// // func Any2Num(o any) float64 {
// // 	v, ok := o.(float64)
// // 	if ok {
// // 		return v
// // 	} else {
// // 		_, ok = o.(Symbol)
// // 		if ok {
// // 			panic("unbound variable" + ShowVal(o))
// // 		} else {
// // 			panic("not number" + ShowVal(o))
// // 		}
// // 	}
// // }

// // func Any2Bool(o any) bool {
// // 	v, ok := o.(bool)
// // 	if ok {
// // 		return v
// // 	} else {
// // 		_, ok = o.(Symbol)
// // 		if ok {
// // 			panic("unbound variable" + ShowVal(o))
// // 		} else {
// // 			panic("not bool" + ShowVal(o))
// // 		}
// // 	}
// // }

// // func cntd(d int) string {
// // 	return fmt.Sprintf("%"+string(rune(2*d))+"s", " ") + string(rune(d)) + " "
// // }

// // func ret(d int, io InOutable, o any) any {
// // 	if d >= 0 {
// // 		io.Out(false, cntd(d)+fmt.Sprintf("%c", 8592))
// // 		io.Out(true, ShowVal(o))
// // 	}
// // 	return o
// // }

// // func trace(d int, io InOutable, o any) {
// // 	if d < 0 {
// // 		return
// // 	}
// // 	io.Out(false, cntd(d)+fmt.Sprintf("%c", 8592))
// // 	io.Out(true, ShowVal(o))
// // }

// // func IsEq(a any, b any) bool {
// // 	if a == nil || b == nil {
// // 		return false
// // 	}
// // 	if reflect.TypeOf(a) != reflect.TypeOf(b) {
// // 		return false
// // 	}
// // 	sa, ok := a.(Symbol)
// // 	if ok {
// // 		sb, _ := b.(Symbol)
// // 		return strings.EqualFold(sa, sb)
// // 	}
// // 	pa, ok := a.(*ConsList)
// // 	if ok {
// // 		pb, _ := b.(*ConsList)
// // 		if pa == pb {
// // 			return true
// // 		}
// // 		ae, be := pa.IsEmpty(), pb.IsEmpty()
// // 		return (ae && be) && (!ae && !be && IsEq(pa.car, pb.car) && IsEq(pa.cdr, pb.cdr))
// // 	}
// // 	ma, ok := a.(*Macros)
// // 	if ok {
// // 		mb, _ := b.(*Macros)
// // 		if ma == mb {
// // 			return true
// // 		}
// // 		return IsEq(ma.pars, mb.pars) && IsEq(ma.body, mb.body)
// // 	}
// // 	return a == b

// // }
func showList(list *List) string {
	builder := new(strings.Builder)

	if !list.IsEmpty() {
		builder.WriteString(showVal(list.Car))
		list = list.Cdr
	}
	for !list.IsEmpty() {
		builder.WriteString(" ")
		builder.WriteString(showVal(list.Car))
		list = list.Cdr
	}
	return builder.String()
}

func showVal(o any) string {
	builder := new(strings.Builder)
	if o == nil {
		builder.WriteString("nil")
	}

	str, ok := o.(string)
	if ok {
		// builder.WriteString("\"")
		builder.WriteString(str)
		// builder.WriteString("\"")
	}
	list, ok := o.(*List)
	if ok {
		builder.WriteString("(")
		builder.WriteString(showList(list))
		builder.WriteString(")")
	}

	if !ok {
		builder.WriteString(fmt.Sprint(o))
	}
	return builder.String()
}
