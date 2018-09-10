package factory

type div struct {
	comput
}

func (d *div) operation() float64 {
	return d.num1 / d.num2
}
