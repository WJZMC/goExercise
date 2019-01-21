package algorithms

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type ChainTabNode struct {
	Data int
	Next *ChainTabNode
	Pre  *ChainTabNode
}

//构造节点
func newChainTabNode(data int) *ChainTabNode {
	node := &ChainTabNode{
		data,
		nil,
		nil,
	}
	return node
}

type ChainTableQueue struct {
	MaxCaps     int
	CurrentCaps int
	Head        *ChainTabNode
	Trail       *ChainTabNode
}

func (q *ChainTableQueue) isEmpty() bool {
	return q.Head == nil
}
func (q *ChainTableQueue) isFull() bool {
	return q.CurrentCaps >= q.MaxCaps
}

//链表队列初始化
func InitChainTableQueue(maxCaps int) *ChainTableQueue {
	queue := &ChainTableQueue{
		maxCaps,
		0,
		nil,
		nil,
	}
	return queue
}

func (q *ChainTableQueue) push(data int) error {
	if q.isFull() {
		return errors.New("队列已满！！！")
	}
	node := newChainTabNode(data)
	if q.Trail == nil || q.Head == nil {
		q.Head = node
		q.Trail = node
	} else {
		node.Pre = q.Trail
		q.Trail.Next = node
		q.Trail = node
	}
	q.CurrentCaps++
	return nil
}

func (q *ChainTableQueue) pop() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("队列为空！！！")
	}
	node := q.Head
	if node.Next == nil {
		q.Head = nil
	} else {
		q.Head = node.Next
	}
	q.CurrentCaps--
	if q.CurrentCaps < 0 {
		q.CurrentCaps = 0
	}
	return node.Data, nil
}
func (q *ChainTableQueue) list() {
	if q.isEmpty() {
		fmt.Println("队列为空！！！")
		return
	}
	var result = "head"
	nodeTmp := q.Head
	for nodeTmp != nil {
		result = fmt.Sprintf("%s==>%d", result, nodeTmp.Data)
		nodeTmp = nodeTmp.Next
	}
	fmt.Println(result)

}

func chainQueue() {
	queue := InitChainTableQueue(5)

	rand.Seed(time.Now().UnixNano())

	mute := sync.RWMutex{}
	go func() {
		for {
			mute.Lock()
			err := queue.push(rand.Intn(100))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
			}
			queue.list()
			mute.Unlock()
			timeral := time.Duration(rand.Int63n(10))
			time.Sleep(timeral * time.Millisecond * 100)
		}
	}()

	for i := 0; i < 2; i++ {
		go func(index int) {
			for {
				mute.RLock()
				value, err := queue.pop()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("=====%d号go程服务%d号客户\n", index, value)
				}
				queue.list()
				mute.RUnlock()
				timeral := time.Duration(rand.Int63n(10))
				time.Sleep(timeral * time.Millisecond * 100)
			}
		}(i)
	}

	for {
		runtime.GC()
	}
}
