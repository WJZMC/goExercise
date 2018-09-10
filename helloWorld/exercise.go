package helloWorld

import (
	"fmt"
)

func HelloWorld() {
	//Exchange()
	//Input()
	//Named()

	//a := -1 >> 2
	//fmt.Println(a)
}

//命名规范
func Named() {

	// 1、字母、数字、 下划线
	// 2、不能数字开头
	// 3、区分大小写
	// 4、不能使用go关键字
	// 5、见名知义
	var int int
	int = 10
	fmt.Println(int)
}

// 格式化输入，输出
func Input() {
	var name string
	var score1, score2 int
	var char byte
	// fmt.Scanf("%s%d%d%d",&name ,&score1 ,&score2 ,&char )
	fmt.Scan(&name, &score1, &score2, &char)

	sum := score1 + score2
	avg := sum / 3
	fmt.Printf("姓名：%s \t 语文：%d \t 数学：%d \t 总分：%d \t 平均分：%d \t code:%c ",
		name, score1, score2,
		sum, avg, char)
}

// 交换值
func Exchange() {
	// var a int = 10
	// var a = 10
	// a := 10
	a, b := 10, 20
	// 方法一
	// c := a
	// a = b
	// b = c
	/*
		方法二
	*/
	// a = a + b
	// b = a - b
	// a = a - b
	/*
		方法三
	*/
	// a, b = b, a
	fmt.Printf("a=%d,b=%d \n", a, b)

}
