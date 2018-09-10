package funcExercise

import "fmt"

type DoubleArgReturnSFunc func(int, int) (int, int)

func Func() {
	//func funcName (int,int)int{}   //func 函数名 （参数） 返回值可选 {}

	//声明、定义、使用
	//sum, sub := funcName1(11, 19)
	//fmt.Println(sum, sub)
	//
	//sum1, sub1 := funcName2(11, 19)
	//fmt.Println(sum1, sub1)

	//funcName3()
	//funcName3(1, 2)
	//funcName3(1, 2, 3, 4)

	//funcTest(1, 2, 3, 4, 5)

	//var f1 = funcTest
	//f1(1, 2, 3, 4)
	//
	////var addSub func(int, int) (int, int)
	//var addSub DoubleArgReturnSFunc
	//addSub = funcName1
	//sum3, sub3 := addSub(11, 19)
	//fmt.Println(sum3, sub3)
	//
	//var f = func() {
	//	fmt.Println("匿名函数")
	//}
	//f()

	//fmt.Printf("%T\n%T,%p,%p", f, funcTest, f, funcTest)
	//f := funcBlock()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Printf("%T,%T,%T", funcBlock, f, f())

}
func funcBlock() func() int {
	var x int
	f := func() int {
		x++
		return x
	}
	return f
}
func funcTest(args ...int) {
	funcName3(args[:]...)
}
func funcName1(a int, b int) (sum int, sub int) { //固定参数和返回值
	sum = a + b
	sub = a - b
	return
}
func funcName2(a, b int) (int, int) { //固定参数和返回值
	sum := a + b
	sub := a - b
	return sum, sub
}
func funcName3(args ...int) { //不定参数可以传0~n个参数
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func Func3(args ...float64) (sum float64, avg float64) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	avg = sum / float64(len(args))
	return //方法一
	// return sum , avg  //方法二

}

//方法三
func Func2(args ...float64) (float64, float64) {
	var sum, avg float64
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	avg = sum / float64(len(args))
	return sum, avg
}
