package goroutine

import (
	"fmt"
	"sync"
)

func SggTest(){

	test1()
}

func test1()  {
	numChan := make(chan int)

	resChan :=make(chan int,2000)

	go insert(numChan)

	//goChanUser :=make(chan bool,8)
	//for i:=0;i<8;i++{
	//	go dealNum(numChan,resChan,goChanUser)
	//
	//}
	//
	//for i:=0;i<8;i++{
	//	<-goChanUser
	//}
	//close(goChanUser)

	wg:=sync.WaitGroup{}
	for i:=0;i<8;i++{
		wg.Add(1)
		go func() {
			for{
				data,ok:=<-numChan
				if !ok{
					break
				}
				result:=0
				for i :=0;i<=data;i++{
					result+=i
				}
				resChan <- result
			}
			wg.Done()
		}()
	}

	wg.Wait()

	for i:=0;i<2000;i++{
		fmt.Println(<-resChan)
	}


}
func insert(numChan chan int)  {
	for i:=0;i<2000;i++{
		numChan<-i
	}
	close(numChan)
}

func dealNum(numChan chan int,resChan chan int,gochan chan bool)  {
	for{
		data,ok:=<-numChan
		if !ok{
			break
		}
		result:=0
		for i :=0;i<=data;i++{
			result+=i
		}
		resChan <- result
	}
	gochan<-true
}