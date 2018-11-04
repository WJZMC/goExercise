package goroutine

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//1、使用channl实现消费者生产者模型。
var condMutex sync.Cond

type production struct {
	id int
}

func question1() {
	ch := make(chan production, 5)
	condMutex.L = new(sync.Mutex)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		go product(ch, i)
	}
	for i := 0; i < 6; i++ {
		go consumer(ch, i)
	}

	for {

	}
}
func product(ch chan<- production, i int) {
	for {
		condMutex.L.Lock()
		for len(ch) == cap(ch) {
			condMutex.Wait()
		}

		object := production{rand.Intn(200)}
		ch <- object
		fmt.Printf("生产者%d,生产货物id=%d\n", i, object.id)
		time.Sleep(time.Second)
		condMutex.Signal()
		condMutex.L.Unlock()

	}
}
func consumer(ch <-chan production, i int) {
	for {
		condMutex.L.Lock()
		for len(ch) == 0 {
			condMutex.Wait()
		}

		obj := <-ch
		fmt.Printf("----------消费者%d,使用了货物id=%d\n", i, obj.id)

		condMutex.Signal()
		condMutex.L.Unlock()
	}
}

//2、假设一个student.txt文件,请你把文件中的信息读出，
// 把每行的数据按age从大到小的顺序写入一个文件student2.txt中。（具体信息放在素材中）
//要求：
//(1) 每行的学生信息可用一个结构体进行保存。
//Type student struct{
//Name string
//Age   int
//}
//(2) 分函数编写。

type student struct {
	name string
	age  int
}

func question2() {
	stuSlice := readFile("/Users/jackwei/Desktop/golang/heima/go进阶/05select 和 锁条件变量/2-其他资料/go提高练习题/素材/student.txt")
	fmt.Println(stuSlice)

	sort(stuSlice)
	fmt.Println(stuSlice)

	writeToFile("/Users/jackwei/Desktop/golang/heima/go进阶/05select 和 锁条件变量/2-其他资料/go提高练习题/素材/student2.txt", stuSlice)

}
func writeToFile(file string, s []student) {
	fp, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	for _, v := range s {
		str := "name=" + v.name + "\t age=" + strconv.Itoa(v.age) + "\n"
		fp.WriteString(str)
	}

}
func readFile(filePath string) []student {

	result := make([]student, 0)

	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		rowStr, err := reader.ReadString('\n')
		part1 := strings.Split(rowStr, "=")
		word := make([]string, 0)
		for _, v := range part1 {
			part2 := strings.Fields(v)
			for _, sub := range part2 {
				word = append(word, sub)
			}
		}
		if len(word) == 4 {
			var stu student
			stu.name = word[1]
			stu.age, _ = strconv.Atoi(word[3])

			result = append(result, stu)
		}
		if err == io.EOF {
			fmt.Println(err)
			break
		}
	}
	return result

}
func sort(s []student) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j].age < s[j+1].age {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
}

//3、从键盘中任意输入一串字符串，编程实现检测字符串中的数字及提取，并安顺序打印到屏幕上。
func question3() {
	reader := bufio.NewReader(os.Stdin)
	strslice, _ := reader.ReadBytes('\n')
	for _, v := range strslice {
		if v >= '0' && v <= '9' {
			fmt.Printf("%c\t", v)
		}
	}

}
