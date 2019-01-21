package algorithms

import "fmt"

type CircleChainTable struct {
	Data int
	Next *CircleChainTable
}

func newCiecleChainNode(data int) *CircleChainTable {
	node := &CircleChainTable{
		data,
		nil,
	}
	return node
}

func (c *CircleChainTable) String() string {
	var result string = fmt.Sprintf("%v", c.Data)

	tmp := c.Next
	for tmp != c {
		result = fmt.Sprintf("%s==>%v", result, tmp.Data)
		tmp = tmp.Next
	}
	return result
}
func InitCircleChain(data int) *CircleChainTable {
	node := newCiecleChainNode(data)
	node.Next = node
	return node
}

func (c *CircleChainTable) insert(data int) {
	tmp := c.Next
	node := newCiecleChainNode(data)

	if tmp == c {
		c.Next = node
		node.Next = c
		return
	}
	last := c
	for tmp != c {
		if tmp.Data >= node.Data {
			break
		}
		last = tmp
		tmp = tmp.Next
	}
	node.Next = tmp
	last.Next = node

}
func (c *CircleChainTable) delete(data int) (head *CircleChainTable) {
	if c.Next == c {
		fmt.Println("空循环链表不能删除")
		return
	}
	cnode, last := c.findCnodeAndLast(data)
	if c == cnode {
		head = last //用来解决删除头节点后选择头节点之后的节点变为头节点
	}
	last.Next = cnode.Next

	return
}
func (c *CircleChainTable) findCnodeAndLast(data int) (cNode, last *CircleChainTable) {
	last = c
	tmp := c.Next
	for tmp != c {
		if tmp.Data == data {
			break
		}
		last = tmp
		tmp = tmp.Next
	}

	return tmp, last
}
func circleChainTable() {
	table := InitCircleChain(0)

	for i := 1; i < 5; i++ {
		if i%2 == 0 {
			table.insert(i * 2)
		} else {
			table.insert(i)
		}
		fmt.Printf("%d==%s\n", i, table)
	}
	table.insert(5)
	fmt.Printf("5==%s\n", table)

	table = table.delete(0)
	fmt.Println("delete 0=", table)

}
