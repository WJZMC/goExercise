package factory

type sub struct {
	comput
}

func (s *sub) operation() float64 {
	return s.num1 - s.num2
}
