package dataType

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func DataType() {
	//Float()
	//Char()
	//String()
	//Fmt()
	//Const()
	//ExerciseDHMS()
	//Operation()
	//Convert()
	//Array()
	//Slice()
	//Map()
	//Struct()

	//a := 10000   //8
	//a := 0.123   //8
	//a := false //1
	//a := '\n' //4 ?

	//a := "我是你发发就发交罚款发发发"//存储的地址，地址是无符号位的整型   //16

	//a := [10]int{1, 2, 3, 4} //80= 8 *10
	//a := []int{1, 2, 3, 6, 7, 87, 8, 8, 8, 8} //24   ?
	//a := map[int]string{1: "woshi", 2: "ofafaf"} //指针 8

	//var a *student //8+16+8+16

	//fmt.Println(unsafe.Sizeof(a))
}

func Float() {
	var a float32 = 1.123456789
	var b float64 = 1.123456789101112131415161718
	fmt.Printf("float32=%f,float64=%f\n", a, b) //%f自动保留6位末位四舍五入
	fmt.Printf("float32=%.7f,float64=%.15f\n", a, b)
	fmt.Printf("float32=%.10f,float64=%.20f\n", a, b) //float32 超过7位出现不准确  float64超过15不准确
}
func String() {
	//b := "I'm jack !"
	//用双引号括起来，单引号括起来默认是字符,字符串结尾默认添加 \0 作为字符串结束标志
	//var c string = " I 我develop iOS 4 years old"
	//var d string = "我"

	//fmt.Println(b)
	//fmt.Printf("%s\n", b+c)
	////一个汉字3个字符，一个字符一个字节
	//fmt.Printf("字符串长度：%d\n%d", len(b+c), len(d))

	//字符串函数

	//result := strings.Contains(c, "I")//是否包含
	//fmt.Printf("%T,%t", result, result)
	//result := strings.Index(c, "develop")//包含索引位置，不包含返回-1：一个汉字三个字符
	//fmt.Printf("%T,%d", result, result)

	//result := strings.Trim(c, " ")//去首位指定字符串
	//fmt.Printf("%T,%s", result, result)

	//result := strings.Repeat(c, 2) //重复
	//fmt.Printf("%T,%s", result, result)
	//result := strings.Replace(c, "e", "你", -1)//字符串替换
	//fmt.Printf("%T,%s", result, result)

	//result := strings.Split(c, " ")//字符串分割
	//fmt.Printf("%T,%v", result, result)
	//result := strings.Fields(c) //按空格分割字符串得到切片
	//fmt.Printf("%T,%s", result, result)

	//slice := strings.Split(c, " ") //字符串分割
	//result := strings.Join(slice, "-")//字符串连接
	//fmt.Printf("%T,%s", result, result)

	fmt.Println(strings.Contains("i am jack","am"))
	fmt.Println(strings.Join([]string{"i","am","jack"}," "))
	fmt.Println(strings.Split("i am jack"," "))
	fmt.Println(strings.Trim(" i am jack "," "))
    fmt.Println(strings.Replace("i am jack","a","#repalce#",-1))
	fmt.Println(strings.Fields("i am jack"))
	fmt.Println(strings.Repeat("i am jack",3))
	fmt.Println(strings.Index("i am jack","am"))
}
func Char() {
	var a byte = 'c' //单引号括起来默认是字符
	fmt.Println(a)
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", a+10)
	var b byte = '\t'
	fmt.Printf("----%c------\n", b)
}
func Const() {
	const MAX = 100
	//常量一般大写、不能改值，
	// 不能左侧赋值
	// 声明的时候必须带值
	// fmt.Printf("%p",&MAX)//error go语言不能获取常量的地址
	fmt.Println(MAX)
	fmt.Printf("%T\n", MAX)
	fmt.Printf("%s", "我是常量，存储在常量区\n") // 直接输出是：字面常量

	// iota枚举  前一个声明iota 后面自动累加
	//1
	// const (
	// 	a = iota
	// 	b
	// 	c
	// 	d
	// )

	// fmt.Println(a, b, c, d)

	// // 2  同一行iota值相同
	// const (
	// 	a = iota
	// 	b = iota
	// 	c = iota
	// 	d = iota
	//  e ,f = iota ,iota
	// )

	//3  iota 不论用不用都会自动累加，可以自定义值，自定义值之后不会自动累加，但是iota会自动累加
	const (
		a = iota
		b = 20
		c = 30
		f
		d = iota
		e
	)
	fmt.Println(a, b, c, d, e, f)

}
func Fmt() {
	var a int = 123
	fmt.Printf("十进制：%d,\t8进制： %o,\t 16进制：%X,\t 16进制：%x\n,\t 类型：%T,\t 地址：%p\n,\t 二进制：%b",
		a, a, a, a, a, &a, a)
	// %t  bool
	// %s 字符串
	// %c 字符
	// %%  %
	// %f  浮点数
	//

}
func ExerciseDHMS() {
	day := 107653

	const offsetM = 60
	const offsetHour = offsetM * offsetM
	const offsetDay = 24 * offsetHour

	d := day / offsetDay
	h := day % offsetDay / offsetHour
	m := day % offsetDay % offsetHour / offsetM
	s := day % offsetDay % offsetHour % offsetM

	fmt.Printf("%d:%d:%d:%d", d, h, m, s)
}

//运算符
func Operation() {

	// // 算术运算
	// a , b := 10 ,20
	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)//被除数不能为0
	// fmt.Println(a%b)//被余数不能为0

	// c , d := 20.1 ,30.1
	// fmt.Println(c+d)
	// fmt.Println(c-d)
	// fmt.Println(c*d)
	// fmt.Println(d/c)//被除数不能为0
	// // fmt.Println(d%c)//error 浮点不能取余，被余数不能为0

	// a ++
	// c --//浮点数可以自加和自减，自加和自减必须单行存在，避免不同操作系统的二义性
	// fmt.Println(a,c)

	// // 赋值运算
	// a = a + b
	// fmt.Println(a)

	// a += b
	// fmt.Println(a)

	// a -= b
	// fmt.Println(a)

	// a *= b
	// fmt.Println(a)

	// a /= b

	// fmt.Println(a)

	// a = a % (b +2)  //如果需要计算 a 取模 b+2  需要给b+2加括号

	// fmt.Println(a)

	//比较运算

	// a , b := 10 ,20
	// fmt.Println(a>b)
	// fmt.Println(a<b)
	// fmt.Println(a<=b)
	// fmt.Println(a>=b)
	// fmt.Println(a!=b)
	// fmt.Println(a==b)

	// 逻辑运算

	// a , b := 10 ,20
	// fmt.Println(a > b || a < b && a == b)// && 优先级大于 ||
	// // fmt.Println(！a )// error 不能对一个非bool类型取！

	//运算符优先级
	/*高
			()
			[]
			.

			! & * ++ --

			* / %
			+ -

			> < >= <= ==

			&&

			||

			= += -+ *= /=

	低*/

}

func Convert() {
	//  基本数据类型（）

	// int()
	// float32()
	// float64()

	// 一般从低到高转
	// 从高到低会损失精度 或者 溢出

	//other ---->string

	//var a int64 = 123
	//str := strconv.FormatInt(a, 8)

	//var a float64 = 1234.56798999
	//str := strconv.FormatFloat(a, 'f', 2, 64)

	//str := strconv.FormatBool(true)

	//var a []byte = []byte{'a', 'b', 'c'}
	//str := string(a)

	//a := 123
	//str := strconv.Itoa(a)

	//str := fmt.Sprintf("%d", a)
	//fmt.Println(str)

	//string ---->other  完全转换，否则就报错

	//str := "1212567"
	//a, error := strconv.Atoi(str)
	//a, error := strconv.ParseInt(str, 10, 64)
	//if error != nil {
	//	fmt.Println(error.Error())
	//} else {
	//	fmt.Printf("%T,%d", a, a)
	//}

	//str := "1234.3455"
	//a, error := strconv.ParseFloat(str, 64)
	//if error != nil {
	//	fmt.Println(error.Error())
	//} else {
	//	fmt.Printf("%T,%f\n", a, a)
	//}

	//str := "true"
	//a, error := strconv.ParseBool(str)
	//if error != nil {
	//	fmt.Println(error.Error())
	//} else {
	//	fmt.Printf("%T,%t\n", a, a)
	//}

	//slice := make([]byte, 0, 100)
	//slice = strconv.AppendBool(slice, true)
	//slice = strconv.AppendFloat(slice, 100.001, 'f', 3, 64)
	//slice = strconv.AppendInt(slice, 123, 8)
	//slice = strconv.AppendQuote(slice, "你是是是是")
	//fmt.Println(string(slice))
}

func Array() {
	//定义
	//var array [4]int = [4]int{0,1,2,3}
	//var array [4]int = [4]int{0,1}  //局部赋值，其他值默认为0
	//var array [4]int = [4]int{1:2}  //指定赋值
	//var array [4]int = [...]int{1:2}

	//array := [4]int{}//类型推导

	//下标赋值
	//array[0]=1

	//	数组取地址需要加 & ，地址指向第一个元素的地址
	//var array [4]int
	//array[0] = 1
	//fmt.Println(array)
}

func Slice() { //下标赋值,空切片为nil，指向0x0，不能直接用下标赋值
	//定义
	//var slice []int = []int{0,1,2,3}
	//var slice []int = []int{0,1}  //局部赋值，其他值默认为0
	//var slice []int = []int{1:2}  //指定赋值
	//下标赋值,空切片为nil，指向0x0，不能直接用下标赋值
	//slice[0]=1//error

	//slice := []int{}//类型推导
	//slice := make([]int,10)//make(类型，长度，容量可选)，长度（0~n）
	//slice=append(slice,1,2)//append(切片，内容不定参列表（1~n)，
	//append如果没有超过容量，地址不发生改变，如果超过容量，切片地址会发生变化

	//切片取地址 不需要加& 直接%P，指向第一个元素的地址，切片截取会影响原来切片的值

	//var slice []int = []int{1, 2, 3, 4, 5}
	//s := make([]int, 5)
	//copy(s, slice)
	//fmt.Println(s)


	//从数组截取  和make（）中的参数不一样
	//slice[low:high:max]
	//low 起始位置
	//high 结束位置
	//len = high - low
	//max : cap = max -low
	//max 跟随最近被截取的切片或者数组
	//arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	//s1:= arr[0:3:9]//s1= [1 2 3] s1长度 3 s1容量 9
	//fmt.Println("s1=",s1,"s1长度",len(s1),"s1容量",cap(s1))
	//s2:= s1[1:5:6]//s2= [2 3 4 5] s1长度 4 s1容量 5
	////s2 虽然是从s1截取，但是实际操作的arr数组，所以high可以截取s1中没有的元素
	////s2 中max <= s1 的max
	//fmt.Println("s2=",s2,"s1长度",len(s2),"s1容量",cap(s2))

	str := []string{"red","","black","","test","","red","black"}
	//str= sliceDel(str,"")
	//sliceDelPoint(&str,"")

	str = delMultKey(str)

	for _,v := range str{
		fmt.Printf("%q",v)
	}

	slice := []int{1,2,3,4,5}
	copytest(&slice)
	fmt.Println(slice[:len(slice)-1])

}
func sliceDel(slice []string,del string)[]string{
	//out := make([]string,0)
	//for _,v:=range slice{
	//	if v != ""{
	//		out = append(out,v)
	//	}
	//}
	for i,v:=range slice{
		if v == del{
			if i+1<=len(slice){
				slice = append(slice[:i],slice[i+1:]...)
			}
		}
	}
	return slice
}
func sliceDelPoint(slice *[]string,del string){
	for i,v:=range *slice{
		if v == del{
			if i+1<=len(*slice){
				*slice = append((*slice)[:i],(*slice)[i+1:]...)
			}
		}
	}
}
//str := []string{"red","","black","","test","","red","black"}

func delMultKey(slice []string)[]string{
	out :=slice[:1]
	for j:= 1;j<len(slice);j++{
		i:=0
		for ; i<len(out);i++{
			if slice[j]==out[i]{
				break
			}
		}
		if i==len(out){
			out= append(out,slice[j])
		}

	}
	return out
}
func copytest(slice *[]int){
	copy((*slice)[2:5],(*slice)[3:5])
}
func Map() { //map 需要初始化后才能赋值，只声明是一个空指针指向0x0

	//定义 var 变量名 map[keytype]valuetype  :
	// keytype可以用= /！=比较的都可以作为key,float,切片，函数不能作为keytype
	//var dic map[int]string = map[int]string{101: "wo", 102: "ta"}
	//dic := map[int]string{101: "ta", 102: "wo"}
	//dic := make(map[int]string)
	//
	//dic[101] = "wo"
	//dic[102] = "ta"
	//fmt.Println(dic)
	//
	////map 需要初始化后才能赋值，只声明是一个空指针指向0x0
	////var dic map[int]string
	////dic[101] = "wo"
	//
	//delete(dic, 101)
	//fmt.Println(dic)

	reader:=bufio.NewReader(os.Stdin)
	str ,_:=reader.ReadString('\n')
	fmt.Println(getMapFromStr(str))

}
func getMapFromStr(str string)map[string]int{
	//slice :=strings.Split(str," ")
	slice :=strings.Fields(str)
	m := make(map[string]int)
	for _,v:=range slice{
		m[v]++
	}
	return m

}

type student struct {
	stid    int
	name    string
	age     int
	address string
}


//封装一个有 string、bool、int、[]string 类型的结构体 Person。
//main函数中声明结构体变量，调用initPerson函数完成结构体赋值，并在main中查看赋值结果
type Person struct {
	name string
	isMan bool
	age int
	interesting []string
}
func (p *Person)PInitWithArg(name string,isMan bool,age int,interesting []string){
	p.name=name
	p.age=age
	p.isMan=isMan
	p.interesting=interesting
}
func PinitWithPersion(p *Person){
	p.name="name"
	p.age=19
	p.isMan=true
	p.interesting=[]string{}
}
func Struct() {
	//var stu student = student{101, "jack", 28, "北京市昌平区"}
	//stu := student{102, "her", 18, ""}
	//var stu student
	//stu.name = "jack"

	//stuSlice := []student{student{101, "jack", 18, "beijing"}, student{102, "her", 18, "changping"}}
	//
	//fmt.Println(stuSlice)

	//结构体作为map的value
	//m := make(map[string]student)
	//m["jack"] = student{1001, "jack", 19, "beijing"}
	//m["tom"] = student{1002, "tom", 18, "beijing"}
	//fmt.Println(m)
	//delete(m, "jack")
	//fmt.Println(m)
	//

	////结构体切片作为map的value
	//sliceM := make(map[string][]student)
	//sliceM["1001"] = []student{student{1001, "jack", 19, "beijing"}, student{1002, "tom", 18, "beijing"}}
	//sliceM["1002"] = []student{student{1001, "jack", 19, "beijing"}, student{1002, "tom", 18, "beijing"}}
	//
	//fmt.Println(sliceM)
	//
	//sliceM["1002"][1].name = "her"
	//fmt.Println(sliceM)
	//
	//sliceM["1002"] = append(sliceM["1002"], student{1003, "test", 100, "beijing"})
	//fmt.Println(sliceM)

	////结构体作为函数参数
	//stu := student{1001, "jack", 18, "beijing"}
	//structTest(stu)
	//fmt.Println(stu)
	//
	////结构体map作为函数参数
	//m := map[string]student{"tom": student{1001, "jack", 18, "beijing"}}
	//structTestMap(m)
	//fmt.Println(m)

	var p Person
	p.PInitWithArg("jack",true,27,[]string{"play games","listing"})
	fmt.Println(p)

	PinitWithPersion(&p)
	fmt.Println(p)



}
func structTest(stu student) {
	stu.age = 19
}
func structTestMap(m map[string]student) {
	temp := m["tom"]
	fmt.Printf("=======%T====%T========", temp, m["tom"])
	temp.age = 20
	m["tom"] = temp
	////结构体作为map的value不能直接用 map[keyType].出来
	////m["tom"].age = 20//error  cannot assign to struct field m["tom"].age in map
	//
	//fmt.Println(m)
}
