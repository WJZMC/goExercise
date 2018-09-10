package pointer

type student struct {
	stid int
	name string
	age  int
}

func Pointer() {

	//指针声明和使用
	//a := 123
	//var p *int = &a

	//var p *int = new(int)
	//*p = 1234
	//fmt.Printf("%T,%p,%v", p, p, *p)

	//数组指针int
	//var arr *[10]int = &[10]int{1, 2, 3, 4, 5, 6}
	//fmt.Printf("%p,%v,%v", arr, *arr, (*arr)[0])

	//指针类型要和接收的指针类型保持一致
	//var p *[10]int
	//arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//p = &arr2
	//fmt.Printf("%T,%p,%T,%p", p, p, arr2, &arr2)

	//指针数组
	//a := 1
	//b := 2
	//var arr [10]*int = [10]*int{&a, &b}
	////arr[0] = &a
	////arr[1] = &b
	//fmt.Printf("%p,%v,%v", &arr, arr, *arr[0])

	////指针变量作为函数参数
	//a, b := 10, 20
	//swap(&a, &b)
	//
	//fmt.Println(a, b)
	//
	////数组指针变量作为函数参数
	//arr := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//bubbleSort(&arr)
	//fmt.Println(arr)

	//指针数组
	//a, b, c := 10, 20, 30
	//arr := [10]*int{&a, &b, &c}
	//fmt.Println(arr)
	//
	//d := [2]int{1, 2}
	//e := [2]int{1, 2}
	//f := [2]int{1, 2}
	//arr2 := [3]*[2]int{&d, &e, &f}
	//fmt.Println(arr2)
	//
	//for _, v := range arr2 {
	//	fmt.Println(*v)
	//}

	//指针和切片
	//slice := []int{9, 3, 42, 1, 43, 4, 5, 546}
	//fmt.Println(slice)
	//p := &slice
	//fmt.Printf("%T,%p,%T,%p\n", p, p, slice, slice)
	////fmt.Println(p[1])//error 指向切片的指针无法根据下标取出对应元素
	//fmt.Println((*p)[1])
	//sliceSort(p)
	//fmt.Println(*p)

	////指针和结构体
	//结构体指针
	//var stu *student = &student{101, "jack", 19, "changping"}
	//fmt.Printf("%p,%v", stu, *stu)
	//stu := student{1001, "jack", 18}
	//structTest(&stu)
	//fmt.Println(stu)

	//结构体数组指针
	//arr := [3]student{student{1001, "jack", 18}, student{1001, "jack", 18}}
	//fmt.Println(arr)
	//p := &arr
	//fmt.Println(*p)
	//p[0].name = "test" //ok
	////(*p)[0].name = "jjj"  //ok
	//fmt.Println(p[0])

	//多级指针
	//a := 10
	//
	//p := &a
	//
	//pp := &p
	//
	//ppp := &pp
	//
	//fmt.Printf("%p,%T\n", ppp, ppp) //0xc42008a020,***int
	//fmt.Printf("%T\n", ***ppp)      //int
	//fmt.Printf("%T\n", **pp)        //int
	//fmt.Printf("%T\n", *p)          //int
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func bubbleSort(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			////常规取值方式
			//if (*arr)[j] > (*arr)[j+1] {
			//	(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			//}
			////或者 ------数组指针可以直接操作数组元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func sliceSort(slice *[]int) {
	//切片使用append函数时，当切面容量发生变化，会改变切片的地址
	*slice = append(*slice, 111, 222, 333, 444)
}
func structTest(stu *student) {
	//结构体变量名，可以简易省略*
	stu.age = 110
	//(*stu).age = 120
}
