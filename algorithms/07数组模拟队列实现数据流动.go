package algorithms

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const ADQMaxSize = 5

type ArrayDataQueue struct {
	Data  [ADQMaxSize]string
	head  int //使用前不需要+1
	trail int //使用前不需要+1，本身就指向下一个待添加位置
}

func InitArrayQueue() *ArrayDataQueue {
	queue := &ArrayDataQueue{
		[ADQMaxSize]string{},
		0,
		0,
	}
	return queue
}
func (q *ArrayDataQueue) isEmpty() bool {
	return q.head == q.trail
}
func (q *ArrayDataQueue) isFull() bool {
	//q.trail+1 表示下下一个索引位置，队列始终空出来一个位置
	return (q.trail+1)%ADQMaxSize == q.head && q.head != q.trail
}

func (q *ArrayDataQueue) size() int {
	return (q.trail + ADQMaxSize - q.head) % ADQMaxSize
}

func (q *ArrayDataQueue) push(data string) error {
	if q.isFull() {
		return errors.New("队列满，请等待！！")
	}
	q.Data[q.trail%ADQMaxSize] = data
	q.trail++
	return nil
}

func (q *ArrayDataQueue) pop() (string, error) {
	if q.isEmpty() {
		return "", errors.New("空队列，不能取数据！！！！")
	}
	value := q.Data[q.head%ADQMaxSize]
	q.head++
	return value, nil
}

func (q *ArrayDataQueue) list() {
	if q.isEmpty() {
		fmt.Println("空队列，不能取数据！！！！")
		return
	}
	tmp := q.head
	result := "队列="
	for tmp != q.trail {
		result = fmt.Sprintf("%s->%s", result, q.Data[tmp%ADQMaxSize])
		tmp++
	}

	fmt.Println(result)
}

func arrayQueueDataStream() {
	//TestSelf()

	queue := InitArrayQueue()

	rand.Seed(time.Now().UnixNano())

	mute := sync.RWMutex{}
	go func() {
		for {
			mute.Lock()
			err := queue.push(strconv.Itoa(rand.Intn(100)))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
			}
			queue.list()
			mute.Unlock()
			timeral := time.Duration(rand.Int63n(100))
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
					fmt.Printf("=====%d号go程服务%s号客户\n", index, value)
				}
				queue.list()
				mute.RUnlock()
				timeral := time.Duration(rand.Int63n(100))
				time.Sleep(timeral * time.Millisecond * 100)
			}
		}(i)
	}

	for {
		runtime.GC()
	}
}

//手动模拟
func TestSelf() {
	queue := InitArrayQueue()

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
			queue.list()

		case "get":
			value, err := queue.pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("获取队列数据：", value)
			}
			queue.list()

		case "show":
			queue.list()
		default:

		}

	}
}
