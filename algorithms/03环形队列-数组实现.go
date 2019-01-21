package algorithms

import (
	"errors"
	"fmt"
)

const (
	CircleCaps = 4
)

//非环形队列，数组实现
type CircleQueue struct {
	value [CircleCaps]string
	head  int //头部，使用不需要+1
	trail int //直接指向即将添加的下标
}

//加入
func (this *CircleQueue) push(value string) error {
	if this.isFull() {
		return errors.New("队列已满,无法添加元素：" + value)
	}
	this.value[this.trail%CircleCaps] = value
	this.trail++

	return nil
}
func (this *CircleQueue) pop() (string, error) {
	if this.isEmpty() {
		return "-1", errors.New("队列为空")
	}
	value := this.value[this.head%CircleCaps]
	this.head++

	return value, nil
}
func (this *CircleQueue) showQueue() {
	if this.isEmpty() {
		fmt.Println("队列为空")
		return
	}
	//循环的时候要加上head，防止出现size小于head后，循环不走
	for i := this.head; i <= this.size()+this.head; i++ {
		fmt.Printf("\tarray[%v]=%v\t", i%CircleCaps, this.value[i%CircleCaps])
	}
	fmt.Println()
	fmt.Printf("size=%v,head=%v,trail=%v,%v\n", this.size(), this.head, this.trail, this.value)

}
func (this *CircleQueue) isEmpty() bool {
	if this.trail == this.head {
		return true
	}
	return false
}
func (this *CircleQueue) isFull() bool {

	//预留一个元素
	//trail本身就指向下一个待添加字段
	//在+1是为了预留一个空间，防止出现首位相接，求长度为0的情况
	if (this.trail+1)%CircleCaps == this.head && this.trail != this.head {
		return true
	}
	return false
}
func (this *CircleQueue) size() int {
	return (this.trail - this.head + CircleCaps) % CircleCaps
}

func circleQueue() {
	queue := &CircleQueue{[CircleCaps]string{}, 0, 0}

	var input string
	var inputValue string
	fmt.Println("输入add添加数据")
	fmt.Println("输入get获取数据")
	fmt.Println("输入show打印所有数据")
	for {
		fmt.Scan(&input)
		switch input {
		case "add":
			fmt.Scan(&inputValue)
			err := queue.push(inputValue)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
			}
			queue.showQueue()

		case "get":
			value, err := queue.pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("获取队列第一个数据：", value)
			}
			queue.showQueue()

		case "show":
			queue.showQueue()
		default:

		}

	}
}
