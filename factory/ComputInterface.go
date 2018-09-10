package factory

import (
	"errors"
	"fmt"
)

type computer interface {
	operation() float64
}
type comput struct {
	num1 float64
	num2 float64
}

func exec(i computer) {
	fmt.Println(i.operation())
}

var (
	UNKNOWN = errors.New("未知运算符")
)

func Comput(num1, num2 float64, operation string) {
	switch operation {
	case "+":
		a := add{comput{num1, num2}}
		exec(&a)
	case "-":
		s := sub{comput{num1, num2}}
		exec(&s)
	case "/":
		d := div{comput{num1, num2}}
		exec(&d)
	default:
		fmt.Println(UNKNOWN)
	}
}
