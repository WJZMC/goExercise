package extision

import "fmt"

func UnitAdd() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)

	return sum
}
func sub(a, b int) int {
	return a - b
}
