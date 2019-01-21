package algorithms

import (
	"errors"
	"fmt"
)

type SingleNode struct {
	Data int
	Next *SingleNode
}

func newNode(data int) *SingleNode {
	node := &SingleNode{
		data,
		nil,
	}
	return node
}

func (this *SingleNode) String() string {
	var result string = "head"
	if this.Next == nil {
		return result
	}
	tmp := this.Next
	for tmp != nil {
		result = fmt.Sprintf("%v--->%v", result, tmp.Data)
		tmp = tmp.Next
	}
	return result
}

func (this *SingleNode) findNodeBefore(oriNode *SingleNode) (*SingleNode, error) {

	if this.Next == nil {
		return nil, errors.New("empty chainTable!!!!!")
	}
	last := this
	tmp := this.Next
	for tmp != nil {
		if tmp.Data >= oriNode.Data {
			break
		}
		last = tmp
		tmp = tmp.Next
	}

	return last, nil
}

func (this *SingleNode) find(data int) (findNode *SingleNode, err error) {
	if this.Next == nil {
		return nil, errors.New("empty chainTable!!!!!")
	}
	tmp := this.Next
	for tmp != nil {
		if tmp.Data == data {
			findNode = tmp
			break
		}
		tmp = tmp.Next
	}

	if findNode == nil {
		return nil, errors.New(" not exist !!!!!")
	}

	return

}

func InitChainTable() *SingleNode {
	return newNode(-1)
}
func (this *SingleNode) insert(data int) {
	node := newNode(data)

	before, err := this.findNodeBefore(node)
	if err != nil {
		this.Next = node
	} else {
		node.Next = before.Next
		before.Next = node
	}

}

func (this *SingleNode) delete(data int) {
	node, err := this.find(data)
	if err == nil {
		before, err := this.findNodeBefore(node)
		if err != nil {
			//fmt.Println(err)
		} else {
			before.Next = node.Next
		}
	}

}
func (this *SingleNode) change(ori, set int) {
	node, err := this.find(ori)
	if err == nil {
		node.Data = set
	}
}
func singleChainTable() {
	head := InitChainTable()

	for i := 1; i < 5; i++ {
		if i%2 == 0 {
			head.insert(i * 2)
		} else {
			head.insert(i)
		}
		fmt.Printf("%d==%s\n", i, head)
	}
	head.insert(5)
	fmt.Printf("5==%s\n", head)

	//
	head.delete(3)
	//
	fmt.Println("delete 3", head)
	//
	head.change(1, 6)
	//
	fmt.Println("change 1 6===", head)
	//
	find, err := head.find(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(find.Data)
	}
	//
	find, err = head.find(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(find.Data)
	}

}
