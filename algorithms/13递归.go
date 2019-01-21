package algorithms

import "fmt"

func test(n int) {
	if n > 0 {
		n--
		test(n)
	}
	fmt.Printf("%v=>\t", n)
}
func recursion() {
	test(10)
}
