package algorithms

import (
	"errors"
	"fmt"
)

const (
	Caps = 4
)

//非环形队列，数组实现
type arrQueue struct {
	value [Caps]string
	head  int //读取端  指向头部的已读上一个元素
	trail int //加入端
}

//加入
func (this *arrQueue) addQueue(value string) error {
	if this.isFull() {
		return errors.New("队列已满,无法添加元素：" + value)
	}
	this.value[this.trail+1] = value
	this.trail++
	return nil
}
func (this *arrQueue) getQueue() (string, error) {
	if this.isEmpty() {
		return "-1", errors.New("队列为空")
	}

	value := this.value[this.head+1]
	this.head++
	return value, nil
}
func (this *arrQueue) show() {
	if this.isEmpty() {
		fmt.Println("队列为空")
		return
	}

	last := this.head
	for this.head != this.trail {
		fmt.Printf("%v ", this.value[this.head+1])
		this.head++
	}
	fmt.Println()
	this.head = last

}
func (this *arrQueue) isEmpty() bool {
	if this.trail == this.head {
		return true
	}
	return false
}
func (this *arrQueue) isFull() bool {
	if this.trail+1 >= Caps {
		return true
	}
	return false
}
func arrayQueue() {

	queue := &arrQueue{[Caps]string{}, -1, -1}

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
			err := queue.addQueue(inputValue)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			value, err := queue.getQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("获取队列第一个数据：", value)
			}
		case "show":
			queue.show()
		default:

		}

	}
}
