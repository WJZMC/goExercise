package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ch = make(chan int)

func Goroutine() {

	//question1()

	//question2()
	//
	//question3()

	//死锁1,同一个go程使用chan同步
	//ch := make(chan int)
	//ch <- 1  //阻塞在这里
	//fmt.Println(<-ch)

	//死锁2,在其他go程开始前使用chan
	//ch := make(chan int)
	//fmt.Println(<-ch) //阻塞在这里
	//go func() {
	//	ch <- 1
	//}()

	//死锁3,多个go程相互等待io操作
	//ch := make(chan int)
	//ch2 := make(chan int)
	//for  {
	//	select {
	//	case data := <-ch:
	//		ch2 <- data
	//	}
	//}
	//for {
	//	select {
	//	case data := <-ch2:
	//		ch <- data
	//	}
	//}

	//逻辑cpu数量模拟
	//count := runtime.NumCPU()
	//fmt.Println(count)
	//runtime.GOMAXPROCS(1)
	//for {
	//	go fmt.Println(0)
	//	fmt.Println(1)
	//}
	//go func() {
	//	i := 0
	//	defer fmt.Println("子go程结束")
	//	for {
	//		fmt.Println("子go程")
	//		i++
	//		if i > 10 {
	//			//runtime.Goexit()
	//		}
	//	}
	//}()
	//
	//for i := 0; i < 50; i++ {
	//	fmt.Println("i=", i)
	//	runtime.Gosched()
	//}

	//runtime.GOMAXPROCS(count)

	//定时器模拟
	//timer := time.NewTimer(time.Second)
	//fmt.Println("now:", time.Now())
	//
	//timer.Reset(time.Second * 2)
	//fmt.Println(<-timer.C)
	//
	//timeTrack := time.NewTicker(time.Second)
	//
	//i := 0
	//for {
	//	fmt.Println(<-timeTrack.C)
	//	i++
	//
	//	if i > 5 {
	//		timeTrack.Stop()
	//		//break
	//		return
	//	}
	//}

	//ch := make(chan order, 6)
	//
	//chPrint := make(chan bool) //处理打印事件，确保同一个订单：先打印出创建订单在输出处理订单
	//go creater(ch, chPrint)
	//dealOrder(ch, chPrint)

	//timer := time.NewTimer(time.Second)
	//fmt.Println(time.Now())
	//timer.Reset(time.Second * 3)
	//fmt.Println(<-timer.C)
	//
	//fmt.Println(<-time.After(time.Second * 4))
	//
	//ticker := time.NewTicker(time.Second * 2)
	//
	//for time := range ticker.C {
	//	fmt.Println(time)
	//}

	//close channel实现
	//ch := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	//for data := range ch {
	//	fmt.Println(data)
	//}
	//ch := make(chan int, 6)
	//
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//for {
	//
	//	if data, ok := <-ch; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//
	//}

	//select应用
	//ch := make(chan int)
	//quit := make(chan bool)
	//go func() {
	//	for {
	//		//num := <-ch
	//		//fmt.Println(num)
	//		select {
	//		case num := <-ch:
	//			fmt.Println(num)
	//		case <-quit:
	//			//goto End
	//			//runtime.Goexit()
	//			return
	//		}
	//	}
	//	//End:
	//}()
	//
	//for i := 0; i < 5; i++ {
	//	ch <- i
	//}
	//quit <- true

	//超时模拟
	//ch := make(chan int)
	//quit := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case num := <-ch:
	//			fmt.Println(num)
	//		case <-time.After(time.Second * 3):
	//			fmt.Println("超时退出")
	//			quit <- true
	//			return
	//		}
	//		fmt.Println("--------------")
	//	}
	//	fmt.Println("-------退出子go程-------")
	//
	//}()
	//
	//for i := 0; i < 5; i++ {
	//	ch <- i
	//	time.Sleep(time.Second * 2)
	//}
	//
	//<-quit

	//斐波那契数列直接实现
	//ch := make(chan int)
	//quit := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case num := <-ch:
	//			fmt.Println(num)
	//		case <-time.After(time.Second * 2):
	//			runtime.Goexit()
	//			quit <- true
	//		}
	//	}
	//}()
	//x, y := 1, 1
	//ch <- x
	//ch <- y
	//for i := 0; i <= (15-2)/2; i++ {
	//	x += y
	//	ch <- x
	//	y += x
	//	ch <- y
	//
	//}
	//<-quit
	//斐波那契数列 递归切片 实现
	//ch := make(chan []int)
	//
	//go func() {
	//	defer fmt.Println("子go程结束")
	//
	//	for {
	//		select {
	//		case result := <-ch:
	//			fmt.Println(result)
	//		case <-time.After(time.Second * 5):
	//			runtime.Goexit()
	//		}
	//	}
	//}()
	//s := []int{1, 1}
	//ch <- s
	//for i := 2; i < 15-2; i++ {
	//	s = getFeibonaqi(s)
	//	ch <- s
	//}

	////超时
	//quit := make(chan bool)
	//ch := make(chan int)
	//go func() {
	//	for {
	//		select {
	//		case num := <-ch:
	//			fmt.Println(num)
	//		case <-time.After(time.Second * 3):
	//			quit <- true
	//			runtime.Goexit()
	//		}
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	ch <- i
	//	time.Sleep(time.Second * 4)
	//}
	//<-quit

	//读写锁 实现访问共享内存
	//seed := time.Now().UnixNano()
	//fmt.Println(seed)
	//rand.Seed(seed)
	//for i := 0; i < 6; i++ {
	//	go readGO(i)
	//}
	//for i := 0; i < 6; i++ {
	//	go writeGO(i)
	//}
	//for {
	//
	//}

	////条件变量  实现生产者消费者模型
	//fmt.Println(time.Now())
	//
	//rand.Seed(time.Now().UnixNano())
	//ch := make(chan order, 7)
	//
	//cond.L = new(sync.Mutex)
	//
	//for i := 0; i < 7; i++ {
	//	go cond_createrOrder(ch)
	//}
	//for i := 0; i < 4; i++ {
	//	go cond_dealOrder(ch)
	//}
	//
	//fmt.Println(time.Now())
	//for {
	//
	//}

	//channel 实现生产者和消费者模型
	//ch := make(chan int, 3)
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for {
	//			for len(ch) != cap(ch) {
	//				data := rand.Intn(200)
	//				ch <- data
	//				fmt.Println("-------w", data)
	//				time.Sleep(time.Second * 2)
	//			}
	//		}
	//	}()
	//}
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for {
	//			for len(ch) != 0 {
	//				fmt.Println("r:", <-ch)
	//			}
	//		}
	//	}()
	//}
	//
	//for {
	//
	//}

}

var cond sync.Cond

func cond_createrOrder(ch chan<- order) {
	for i := 0; i < 100; i++ {
		cond.L.Lock()
		for len(ch) == cap(ch) {
			cond.Wait()
		}
		o := order{rand.Intn(200)}
		ch <- o
		fmt.Println("创建订单id：", o.id)
		cond.Signal()
		cond.L.Unlock()
	}
}
func cond_dealOrder(ch <-chan order) {
	for {
		cond.L.Lock()
		for len(ch) == 0 {
			cond.Wait()
		}
		o := <-ch
		fmt.Println("处理订单：", o.id)
		cond.Signal()
		cond.L.Unlock()
	}
}

var rwMutex sync.RWMutex
var share int

func readGO(index int) {
	for {
		rwMutex.RLock()
		fmt.Println(index, "\t r:", share)
		time.Sleep(time.Second)
		rwMutex.RUnlock()
	}
}
func writeGO(index int) {
	for {
		rwMutex.Lock()
		data := rand.Intn(200)
		share = data
		fmt.Println(index, "\t w:", data)
		time.Sleep(time.Second)
		rwMutex.Unlock()
	}
}

func getFeibonaqi(s []int) []int {
	return append(s, s[len(s)-2]+s[len(s)-1])
}

type order struct {
	id int
}

func creater(ch chan<- order, chPrint chan bool) {
	for i := 0; i < 100; i++ {
		ord := order{i}
		ch <- ord
		fmt.Println("创建订单：", ord.id)
		chPrint <- true
	}
	close(ch)
}
func dealOrder(ch <-chan order, chPrint chan bool) {
	//for {
	//	if data, ok := <-ch; ok {
	//		<-chPrint
	//		fmt.Println("处理订单：", data)
	//	} else {
	//		break
	//	}
	//}
	for data := range ch {
		<-chPrint
		fmt.Println("订单处理：", data)
	}
}

var mutex sync.Mutex

func go1() {
	printf("first goroutine")
	//ch <- 1

}
func go2() {
	//<-ch
	printf("sec goroutine")
}

func printf(s string) {
	mutex.Lock()
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}
