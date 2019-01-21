package algorithms

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	output []int
)

type JosephulNode struct {
	data int
	Next *JosephulNode
}

func InitJosePhlNode(data int) *JosephulNode {
	node := &JosephulNode{
		data,
		nil,
	}
	return node
}
func (j *JosephulNode) isEmpyt() bool {
	return j.Next == j
}
func (j *JosephulNode) append(data int) error {
	trail := j
	for trail.Next != j {
		trail = trail.Next
	}

	node := InitJosePhlNode(data)
	trail.Next = node
	node.Next = j
	return nil
}
func (j *JosephulNode) delete(data int) (*JosephulNode, error) {
	if j.isEmpyt() {
		return j, errors.New("empty !!!!")
	}
	current, last, err := j.findCurrentAndLast(data)
	if err == nil {
		output = append(output, current.data)
		last.Next = current.Next
		return current.Next, nil
	}
	return j, err
}
func (j *JosephulNode) findCurrentAndLast(data int) (curent, last *JosephulNode, err error) {
	last = j
	tmp := j.Next
	for tmp != j {
		if tmp.data == data {
			curent = tmp
			break
		}
		last = tmp
		tmp = tmp.Next
	}

	if curent == nil {
		return nil, nil, errors.New("not exist!!!!")
	}

	return curent, last, nil
}

func (j *JosephulNode) list() {
	var result = "head=" + strconv.Itoa(j.data)

	tmp := j.Next
	for tmp != j {
		result = fmt.Sprintf("%v==>%d", result, tmp.data)
		tmp = tmp.Next
	}

	fmt.Println(result)
}

func (j *JosephulNode) mapChainTable(index int) *JosephulNode {
	tmp := j
	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}
	tmp, err := j.delete(tmp.data)
	if err == nil {
		return tmp
	}
	return j
}

func JosephulTest() {
	josephul := InitJosePhlNode(1)
	josephul.Next = josephul

	n := 10
	for i := 2; i <= n; i++ {
		josephul.append(i)
		josephul.list()
	}

	//josephul, _ = josephul.delete(5)
	//josephul.list()

	rand.Seed(time.Now().UnixNano())
	tmpCaps := n
	tmpJose := josephul
	for !tmpJose.isEmpyt() {
		deleteNum := rand.Intn(tmpCaps) + 1
		fmt.Println(deleteNum)
		tmpJose = tmpJose.mapChainTable(deleteNum)
		tmpCaps--
	}

	output = append(output, tmpJose.data)
	fmt.Println("出队序列：", output)

}
