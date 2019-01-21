package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

const LengthT = 100000

func sort() {
	rand.Seed(time.Now().UnixNano())

	var array1 [LengthT]int
	for i := 0; i < LengthT; i++ {
		array1[i] = rand.Intn(LengthT)
	}

	//fmt.Println(array1)
	fmt.Println("begin--1---", time.Now().Unix())
	BubbleSort(&array1)
	//selectSort(&array)
	//insertSort(&array)
	//quickSort(&array, 0, len(array)-1)
	fmt.Println("end---1--", time.Now().Unix())
	//fmt.Println(array1)

	var array2 [LengthT]int
	for i := 0; i < LengthT; i++ {
		array2[i] = rand.Intn(LengthT)
	}

	//fmt.Println(array2)
	fmt.Println("begin---2--", time.Now().Unix())
	//BubbleSort(&array2)
	selectSort(&array2)
	//insertSort(&array)
	//quickSort(&array, 0, len(array)-1)
	fmt.Println("end--2---", time.Now().Unix())
	//fmt.Println(array2)

	var array3 [LengthT]int
	for i := 0; i < LengthT; i++ {
		array3[i] = rand.Intn(LengthT)
	}

	//fmt.Println(array3)
	fmt.Println("begin---3--", time.Now().Unix())
	//BubbleSort(&array3)
	//selectSort(&array)
	insertSort(&array3)
	//quickSort(&array, 0, len(array)-1)
	fmt.Println("end----3-", time.Now().Unix())
	//fmt.Println(array3)

	var array4 [LengthT]int
	for i := 0; i < LengthT; i++ {
		array4[i] = rand.Intn(LengthT * 2)
	}

	fmt.Println(array4)
	fmt.Println("begin--4---", time.Now().Unix())
	//BubbleSort(&array4)
	//selectSort(&array)
	//insertSort(&array)
	quickSort(&array4, 0, len(array4)-1)
	fmt.Println("end---4--", time.Now().Unix())
}

func BubbleSort(array *[LengthT]int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if (*array)[j] < (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
}
func selectSort(array *[LengthT]int) {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if (*array)[min] > (*array)[j] {
				min = j
			}
		}
		if i != min {
			(*array)[i], (*array)[min] = (*array)[min], (*array)[i]
		}
	}
}
func insertSort(array *[LengthT]int) {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if (*array)[i] <= (*array)[j] {
				(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
			}
		}
	}
}

func quickSort(array *[LengthT]int, begin, end int) {
	l := begin
	r := end
	base := (*array)[l] //假设取数组第一个元素作为base
	for l < r {
		for (*array)[r] > base { //从右往左找比base大的，指针左移
			r--
		}

		for (*array)[l] < base { //从左往右找比base小的，指针右移
			l++
		}
		//--和++后需要判断，满足才能交换
		if l >= r {
			break
		}
		//从右往左找到比base小的，和从左往右找到比base大的，交换位置
		(*array)[l], (*array)[r] = (*array)[r], (*array)[l]
		//如果有想等元素会一直死循环
		if (*array)[l] == (*array)[r] {
			l++
			r--
		}
	}
	//如果想等，则不需要包含l和r，因为左边小于，右边大于，包含l和r也没有意义
	//如果想等，不处理移动的话，会形成死递归
	if l == r {
		l++
		r--
	}
	if begin < r {
		quickSort(array, begin, r)
	}
	if l < end {
		quickSort(array, l, end)
	}

}

//2 4 1 5 8
//l 0
//r 4
//base 2
//从右往左，找到比base小的1 时 r=2，将array[r]赋值给array[l]
//1 4 1 5 8
//从左往右，找到比base大的4 时 l=1，将array[l]赋值给array[r]
//1 4 4 5 8
//将base赋值给array[l]
//1 2 4 5 8

//再从右往左  r=1 将array[r]赋值给array[l]
//1 2 4 5 8
//从左往右，找到比base大的4 时 l=2，将array[l]赋值给array[r]
//判断是否满足，不满足退出
