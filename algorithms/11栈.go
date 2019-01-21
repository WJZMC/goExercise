package algorithms

import (
	"errors"
	"fmt"
)

const MaxStackCaps = 5

type Stack struct {
	Value [MaxStackCaps]int
	Top   int
}

func NewStack() *Stack {
	return &Stack{
		[MaxStackCaps]int{},
		-1,
	}
}
func (s *Stack) push(data int) error {
	if s.Top > MaxStackCaps-1 {
		return errors.New("full!!!")
	}
	s.Top++
	s.Value[s.Top] = data
	return nil
}
func (s *Stack) pop() (int, error) {
	if len(s.Value) == 0 || s.Top == -1 {
		return -1, errors.New("empty!!!")
	}
	value := s.Value[s.Top]
	s.Top--
	return value, nil
}
func (s *Stack) list() {
	for i := s.Top; i > -1; i-- {
		fmt.Printf("%v\t", s.Value[i])
	}
	fmt.Println()
}

func stackTest() {
	stack := NewStack()
	for i := 0; i < 6; i++ {
		err := stack.push(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("add sucess!!")
		}
	}
	value, err := stack.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)

	stack.list()

}
