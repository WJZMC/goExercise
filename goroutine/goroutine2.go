package goroutine

import (
	"sync"
	"fmt"
	"time"
)

func GetPrime() {
	//forTest()
	//lockTest()
	chanTest()

}

//chan
func chanTest()  {
	ct1 :=time.Now().Unix()

	oriCh :=make(chan int,100000)

	gropuChan :=make(chan bool,8)

	for i:=0;i<8;i++{
		go dealPrima(oriCh,gropuChan)
	}

	for i := 2; i < 100000; i++ {
		oriCh<-i
	}
	close(oriCh)

	for i:=0;i<8;i++{
		<-gropuChan
	}


	ct2 :=time.Now().Unix()

	fmt.Println(ct1,"\n",ct2)

}
func dealPrima(orich chan int,groupCh chan bool)  {
	for{
		tmp , ok := <-orich
		if !ok{
			break
		}
		result:=true
		for j := 2; j< tmp; j++ {
			if tmp%j == 0 {
				result=false
				break
			}
		}
		if result{
			fmt.Println(tmp)
		}
	}
	groupCh<-true
}
//mutex
func lockTest()  {
	slice :=make([]int,0)
	ct1 :=time.Now().Unix()


	lock:=sync.Mutex{}

	group := sync.WaitGroup{}
	for i := 2; i < 100000; i++ {
		group.Add(1)
		go func(index int) {
			result:=true
			tmp:=index
			for j := 2; j< tmp; j++ {
				if tmp%j == 0 {
					result=false
					break
				}
			}
			if result{
				//fmt.Println(i)
				lock.Lock()
				slice=append(slice, tmp)
				lock.Unlock()
			}
			group.Done()
		}(i)
	}

	group.Wait()

	ct2 :=time.Now().Unix()

	fmt.Println(ct1,"\n",ct2)
	lock.Lock()
	for _,v:=range slice{
		fmt.Println(v)
	}
	lock.Unlock()
}
//for
func forTest()  {
	slice :=make([]int,0)
	ct1 :=time.Now().Unix()

	for i := 2; i < 20000; i++ {
		result:=true
		for j := 2; j< i; j++ {
			if i%j == 0 {
				result=false
				//break
			}
		}
		if result{
			slice=append(slice, i)
		}
	}
	ct2 :=time.Now().Unix()

	fmt.Println(ct1,"\n",ct2)
	for _,v:=range slice{
		fmt.Println(v)
	}
}
