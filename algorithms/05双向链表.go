package algorithms

import (
	"errors"
	"fmt"
)

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Pre  *DoubleNode
}

func newDoubleNode(data int) *DoubleNode {
	node := &DoubleNode{
		data,
		nil,
		nil,
	}
	return node
}

func (this *DoubleNode) String() string {
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

func (this *DoubleNode) find(data int) (findNode *DoubleNode, err error) {
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

func InitDoubleChainTable() *DoubleNode {
	return newDoubleNode(-1)
}
func (this *DoubleNode) insert(data int) {
	node := newDoubleNode(data)

	//查找插入位置
	last := this
	tmp := this.Next
	for tmp != nil {
		if tmp.Data >= node.Data {
			break
		}
		last = tmp
		tmp = tmp.Next
	}
	node.Pre = last
	last.Next = node
	if tmp == nil {
		node.Next = nil
		return
	}
	node.Next = tmp
	tmp.Pre = node

}

func (this *DoubleNode) delete(data int) {
	node, err := this.find(data)
	if err == nil {
		before := node.Pre
		after := node.Next
		if before == nil {
			return
		}
		if after == nil {
			before.Next = nil
			return
		}
		before.Next = after
		after.Pre = before
	}

}
func (this *DoubleNode) change(ori, set int) {
	node, err := this.find(ori)
	if err == nil {
		node.Data = set
	}
}
func DoubleChainTable() {
	head := InitDoubleChainTable()

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

	head.delete(3)

	fmt.Println("delete 3", head)

	head.change(1, 6)

	fmt.Println("change 1 6===", head)

	find, err := head.find(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(find.Data)
	}

	find, err = head.find(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(find.Data)
	}

}
