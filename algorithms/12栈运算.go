package algorithms

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	MaxNumStack       = 10
	MaxOperationStack = 10
)

type numStack struct {
	array [MaxNumStack]int
	top   int
}

type operationStack struct {
	array [MaxOperationStack]byte
	top   int
}

func (n *numStack) push(data int) error {
	if n.top >= MaxNumStack {
		return errors.New("full")
	}
	n.top++
	n.array[n.top] = data
	return nil
}
func (o *operationStack) push(operatin byte) error {
	if o.top >= MaxOperationStack {
		return errors.New("full")
	}
	o.top++
	o.array[o.top] = operatin
	return nil
}

func (n *numStack) pop() (int, error) {
	if n.top < 0 {
		return -1, errors.New("empty")
	}
	value := n.array[n.top]
	n.top--
	return value, nil
}
func (o *operationStack) pop() (byte, error) {
	if o.top < 0 {
		return 0, errors.New("empty")
	}
	value := o.array[o.top]
	o.top--
	return value, nil
}
func (s *operationStack) list() {
	for i := s.top; i > -1; i-- {
		fmt.Printf("%v\t", s.array[i])
	}
	fmt.Println()
}
func (s *numStack) list() {
	for i := s.top; i > -1; i-- {
		fmt.Printf("%v\t", s.array[i])
	}
	fmt.Println()
}

func getNumArray(test string) []int {
	reg := regexp.MustCompile(`[0-9]*`)
	result := reg.FindAllStringSubmatch(test, -1)
	fmt.Println(result)
	var num []int
	for i := 0; i < len(result); i++ {
		tmp, _ := strconv.Atoi(result[i][0])
		num = append(num, tmp)
	}
	return num
}
func getOpt(test string) []byte {
	reg := regexp.MustCompile(`[^0-9]`)
	result := reg.FindAllStringSubmatch(test, -1)
	fmt.Println(result)
	var opt []byte
	for i := 0; i < len(result); i++ {
		tmp := []byte(result[i][0])
		opt = append(opt, tmp[0])
	}
	return opt
}
func getLevel(opt byte) (result int) {

	switch opt {
	case '+', '-':
		result = 0
	case '*', '/':
		result = 1
	default:
		result = 0
	}
	return
}
func cal(num1, num2 int, opt byte) (result int) {
	switch opt {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	default:
		result = 0
	}
	return
}

func stackOperation() {
	test := "32+2*4-10/2"

	numArray := getNumArray(test)
	optArray := getOpt(test)
	fmt.Println(optArray)

	numIndex, optIndex := 0, 0

	nums := &numStack{[MaxNumStack]int{}, -1}
	optS := &operationStack{[MaxOperationStack]byte{}, -1}

	nums.push(numArray[numIndex])
	optS.push(optArray[optIndex])
	numIndex++
	nums.push(numArray[numIndex])
	for {
		optTmp, err := optS.pop()
		fmt.Println("optS.pop()---1", optTmp)
		if err != nil {
			optS.push(optArray[optIndex])
			fmt.Println("optS.push(optArray[optIndex])---2", optTmp)
		} else {
			optIndex++
			if optIndex > len(optArray)-1 {
				optS.push(optTmp)
				break
			}
			optNext := optArray[optIndex]
			if getLevel(optTmp) > getLevel(optNext) {
				numStaPopAfter, err := nums.pop()
				if err != nil {
					break
				}
				numStaPop, err := nums.pop()
				if err != nil {
					break
				}
				result := cal(numStaPop, numStaPopAfter, optTmp)
				nums.push(result)
				fmt.Println("nums.push(result)---3", result)
				optIndex--

			} else {
				optS.push(optTmp)
				optS.push(optNext)
				fmt.Println("optS.push(optTmp)---4", optTmp, optNext)
				numIndex++
				if numIndex > len(numArray)-1 {
					break
				}
				nums.push(numArray[numIndex])

			}
		}
		nums.list()
		optS.list()
		fmt.Println("numindex=5", numIndex)
		fmt.Println("optindex=6", optIndex)

	}
	nums.list()
	optS.list()
	for {
		optTmp, err := optS.pop()
		fmt.Println("optS.pop()---7", optTmp)
		if err != nil {
			break
		} else {
			numStaPopAfter, err := nums.pop()
			if err != nil {
				break
			}
			numStaPop, err := nums.pop()
			if err != nil {
				break
			}
			result := cal(numStaPop, numStaPopAfter, optTmp)
			nums.push(result)
		}
	}
	numStaPop, err := nums.pop()
	if err != nil {
		fmt.Println("-1::", err)
		return
	}
	fmt.Println("end-----", numStaPop)

}
