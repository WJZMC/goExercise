package factory

type add struct {
	comput
}

func (a *add) operation() float64 {
	return a.num1 + a.num2
}
