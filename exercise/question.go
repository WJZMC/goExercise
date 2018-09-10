package exercise

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Exercise() {
	//Question1()
	//Question2()
	//Question3()
	//Question4()
	//exercise.Question5()
	//Question6()
	//exercise.Question7()
	//Question8()
	//Question9()//键盘输入文本带空格接收问题不清晰
	//Question10()
	//Question11()
	//Question12()
	//Question13()
	//Question14()//字符串和字符数组互转不清晰
	//Question15()
	//Question16()
	//Question17()
	//Question18() // 切片需要学习
	//Question19()
	//Question20()
	//Question21()
	//Question22()
	//Question23()
	//Question24()
	//Question25()
	//Question26()
	//Question27()
	//Question28()
	//Question29()
	//Question30()
}

// 1、题目：有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
// 程序分析：可填在百位、十位、个位的数字都是1、2、3、4。组成所有的排列后再去掉不满足条件的排列。
func Question1() {
	var count int
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 4; k++ {
				if i != j && i != k && j != k {
					count++
					fmt.Printf("%d\t", i*100+j*10+k)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(count)

}

// 2、题目：输入某年某月某日，判断这一天是这一年的第几天？
// 程序分析：以1987年3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，
// 特殊情况，闰年且输入月份大于3时需考虑多加一天。
func Question2() {
	for {
		var year, month, day, curentDay int
		fmt.Scan(&year, &month, &day)
		//效率低
		//for i := 1; i < month; i++ {
		//	switch i {
		//	case 1, 3, 5, 7, 8, 10, 12:
		//		curentDay += 31
		//	case 2:
		//		if (year%4 == 0 && year%100 != 0 || year%400 == 0) && month > 2 {
		//			curentDay += 29
		//		} else {
		//			curentDay += 28
		//		}
		//	default:
		//		curentDay += 30
		//	}
		//}
		//高效
		//switch 月份倒排，用fallthrough
		switch month - 1 {
		case 11:
			curentDay += 30
			fallthrough
		case 10:
			curentDay += 31
			fallthrough
		case 9:
			curentDay += 30
			fallthrough
		case 8:
			curentDay += 31
			fallthrough
		case 7:
			curentDay += 31
			fallthrough
		case 6:
			curentDay += 30
			fallthrough
		case 5:
			curentDay += 31
			fallthrough
		case 4:
			curentDay += 30
			fallthrough
		case 3:
			curentDay += 31
			fallthrough
		case 2:
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				curentDay += 29
			} else {
				curentDay += 28
			}
			fallthrough
		case 1:
			curentDay += 31
		default:
			curentDay += 0
		}

		curentDay += day
		fmt.Printf("这一天是这一年的第%d天\n", curentDay)
	}
}

// 3、题目：输入三个整数x,y,z，请把这三个数由小到大输出。
// 程序分析：我们想办法把最小的数放到x上，先将x与y进行比较，如果x>y则将x与y的值进行交换，
// 然后再用x与z进行比较，如果x>z则将x与z的值进行交换，这样能使x最小。
func Question3() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Println(x, y, z)
}

// 4、题目：输出9*9口诀。
// 程序分析：分行与列考虑，共9行9列，i控制行，j控制列。
func Question4() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
}

// 5、题目：判断101-200之间有多少个素数，并输出所有素数。
// 程序分析：判断素数的方法：用一个数分别去除2到这个数，如果能被整除，则表明此数不是素数，反之是素数。
func Question5() {
	var count int
	for i := 101; i <= 200; i += 2 {
		isSu := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isSu = false
				break
			}
		}
		if isSu {
			count++
			// fmt.Println(i)
		}
	}
	fmt.Println(count)
}

// 6、题目：打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。
// 例如：153是一个“水仙花数”，因为153=1的三次方＋5的三次方＋3的三次方。
// 程序分析：利用for循环控制100-999个数，每个数分解出个位，十位，百位。
func Question6() {
	var count int
	for i := 100; i <= 999; i++ {
		xxx, xx, x := i/100, i%100/10, i%10
		if xxx*xxx*xxx+xx*xx*xx+x*x*x == i {
			fmt.Println(xxx, "\t", xx, "\t", x, "\t")
			count++
		}
	}
	fmt.Println(count)
}

// 7、题目：将一个正整数分解质因数。例如：输入90,打印出90=2*3*3*5。
// 程序分析：对n进行分解质因数，应先找到一个最小的质数k，然后按下述步骤完成：
// (1)如果这个质数恰等于n，则说明分解质因数的过程已经结束，打印出即可。
// (2)如果n<>k，但n能被k整除，则应打印出k的值，并用n除以k的商,作为新的正整数你n,重复执行第一步。
// (3)如果n不能被k整除，则用k+1作为k的值,重复执行第一步。
func Question7() {
	var n int
	for {
		fmt.Scan(&n)
		if n <= 2 {
			fmt.Println("请输入大于2的数字")
			continue
		}
		result := fmt.Sprintf("%d=", n)
		// fmt.Println(result,"\t",n)
		question7_func(n, 2, result)
	}
}
func question7_func(n, i int, result string) {
	if i < n {
		if n%i == 0 {
			result += fmt.Sprintf("%d*", i)
			n /= i
		} else {
			i++
		}
		question7_func(n, i, result)
	} else if i == n {
		result += fmt.Sprintf("%d", i)
		fmt.Println(result)
	} else {
		fmt.Println(result)
	}

}

// 8、题目：输入两个正整数m和n，求其最大公约数和最小公倍数。
// 程序分析：利用辗除法。
func Question8() {
	var m, n int
	for {
		fmt.Scan(&m, &n)
		if m > 0 && n > 0 {
			if m > n {
				m, n = n, m
			}
			a, b := m, n
			question8_func(a, b, m, n)
		} else {
			fmt.Println("请输入正整数！")
		}
	}
}
func question8_func(a, b, m, n int) {
	// fmt.Println(a ,b )
	if b%a != 0 {
		a = b % a
		b = m
		//    fmt.Println(a ,b )
		question8_func(a, b, m, n)
	} else {
		fmt.Println("最大公约数：", a, "最小公倍数：", m*n/a)
	}
}

// 9、题目：输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
func Question9() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
	English, space, num, other := 0, 0, 0, 0
	for i := 0; i < len(str); i++ {
		temp := str[i]
		// fmt.Println(temp)
		if temp >= 48 && temp <= 57 {
			num++
		} else if temp >= 65 && temp <= 90 || temp >= 97 && temp <= 122 {
			English++
		} else if temp == 32 {
			space++
		} else {
			other++
		}
	}
	fmt.Printf("字母:%d \t 空格：%d \t 数字:%d \t 其他字符:%d \n", English, space, num, other)
}

// 10、题目：一球从100米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第10次落地时，
// 共经过多少米？第10次反弹多高（使用float64 保留两位小数）？
func Question10() {
	height := 100.0
	maxH := 0.0
	for i := 0; i < 10; i++ {
		if i != 0 {
			maxH += height * 2
		} else {
			maxH += height
		}
		height /= 2
		// fmt.Println(height,maxH)
	}
	fmt.Printf("共经过%.2fm,第10次%.2fm\n", maxH, height)
}

// 11、题目：打印出如下图案（菱形*）1，3，5，7，5，3，1
// 程序分析：先把图形分成两部分来看待，前四行一个规律，后三行一个规律，利用双重for循环，第一层控制行，第二层控制列。
func Question11() {
	for {
		fmt.Print("请输入奇数行数：")
		var row int
		fmt.Scan(&row)

		//前后空格规律
		// mid := row  / 2
		// for i := 0 ;i < row ; i++{
		// 	for j := 0 ;j < row ;j ++{
		// 		if i <= mid {
		// 			if j < mid - i || j >= row -(mid - i) {
		// 				// fmt.Print(i,j)
		// 				fmt.Print(" ")
		// 			}else{
		// 				fmt.Print("*")
		// 			}
		// 		}else{
		// 			if j < i - mid || j >= row -( i - mid ) {
		// 				// fmt.Print(i,j)
		// 				fmt.Print(" ")
		// 			}else{
		// 				fmt.Print("*")
		// 			}
		// 		}
		// 	}
		// 	fmt.Println()
		// }

		//*规律
		mid := row / 2
		for i := 0; i < row; i++ {
			for j := 0; j < row; j++ {
				if i <= mid {
					beginIndex := mid - (2*i+1)/2
					endIndex := beginIndex + 2*i + 1
					if j >= beginIndex && j < endIndex {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				} else {
					beginIndex := (2*i+1)/2 - mid
					endIndex := row - beginIndex
					if j >= beginIndex && j < endIndex {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}

		// 算空格
		// for i := 1 ; i <= row ; i ++{
		// 	for j := 1 ;j <= 2 * row - 1 ; j++{
		// 		if j<= row - i || j > 2 * row - 1 -(row -i) {
		// 			fmt.Print(" ")
		// 		}else{
		// 			fmt.Print("*")
		// 		}
		// 	}
		// 	fmt.Println()
		// }

		// 算*号
		//for i := 1; i <= row; i++ {
		//	for j := 1; j <= 2*row-1; j++ {
		//		beginIndex := row - (2*i-1)/2
		//		endIndex := beginIndex + 2*i - 1
		//		// fmt.Print(beginIndex,endIndex)
		//		if j >= beginIndex && j < endIndex {
		//			fmt.Print("*")
		//		} else {
		//			fmt.Print(" ")
		//		}
		//	}
		//	fmt.Println()
		//}
	}
}

// 12、题目：有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的前20项之和。
// 程序分析：请抓住分子与分母的变化规律。（使用float64 保留两位小数）
func Question12() {
	a, b, sum := 2.0, 1.0, 0.0
	for i := 1; i < 20; i++ {
		sum += a / b
		a, b = b, a
		a += b
		fmt.Println(a, b)
	}
	fmt.Printf("%.2f", sum)
}

// 13、题目：求1+2!+3!+...+20!的和
// 程序分析：此程序只是把累加变成了累乘。
func Question13() {
	sum, temp := 0, 1
	for i := 1; i <= 20; i++ {
		temp = 1
		for j := 1; j <= i; j++ {
			temp *= j
		}
		sum += temp
	}
	fmt.Println(sum)
}

// 14、题目：利用递归函数调用方式，将所输入的5个字符，以相反顺序打印出来。
func Question14() {
	var str string
	fmt.Scan(&str)
	question14_func([]byte(str), 0)
}
func question14_func(array []byte, index int) {
	if index < len(array)/2 {
		array[index], array[len(array)-index-1] = array[len(array)-index-1], array[index]
		index++
		question14_func(array, index)
	} else {
		fmt.Println(string(array[:]))
	}
}

// 15、题目：有5个人坐在一起，问第五个人多少岁？他说比第4个人大2岁。问第4个人岁数，他说比第3个人大2岁。
// 问第三个人，又说比第2人大两岁。问第2个人，说比第一个人大两岁。最后问第一个人，他说是10岁。请问第五个人多大？
// 程序分析：利用递归的方法，递归分为回推和递推两个阶段。要想知道第五个人岁数，
// 需知道第四人的岁数，依次类推，推到第一人（10岁），再往回推。
func Question15() {
	var count int
	fmt.Scan(&count)
	question15_func(count, 10)
}
func question15_func(count, age int) {
	if count > 1 {
		age += 2
		count--
		question15_func(count, age)
	} else {
		fmt.Println(age)
	}
}

// 16、题目：对10个数进行排序
// 程序分析：可以利用选择法，即从后9个比较过程中，选择一个最小的与第一个元素交换，
// 下次类推，即用第二个元素与后8个进行比较，并进行交换。
func Question16() {

	// 冒泡 小--》大  运用：切片

	//slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//sortBubble(slice)
	//fmt.Println(slice)

	//// 选择  小--》大   运用：函数返回值
	//arr := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//arr = sortSelect(arr)
	//fmt.Println(arr)

	//插入  小--》大    运用：数组指针

	arr := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sortInsert(&arr)
	fmt.Println(arr)

}
func sortSelect(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
		fmt.Println(arr)
	}
	return arr
}
func sortInsert(arr *[10]int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if (*arr)[j-1] > (*arr)[j] {
				(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
			}
		}
		//fmt.Println(arr)
	}
	return
}
func sortBubble(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j+1], slice[j] = slice[j], slice[j+1]
			}
		}
		//fmt.Println(slice)
	}
}

// 17、题目：打印出杨辉三角形（要求打印出10行如下图）
// 程序分析：
// 			 1
// 			1 　1
// 			1 　2 　1
// 			1　 3 　3　 1
// 			1　 4　 6 　4 　1
// 			1　 5　 10　10　5 　1
func Question17() {
	var row int = 10
	// fmt.Print("请输入行数：")
	// fmt.Scan(&row)
	var array [10][10]int
	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				array[i][j] = 1
				fmt.Printf("1\t")
				// fmt.Println(i,j)
			} else {
				array[i][j] = array[i-1][j-1] + array[i-1][j]
				fmt.Print(array[i][j], "\t")
			}
		}
		fmt.Println()
	}
}

// 18、题目：有n个人围成一圈，顺序排号。从第一个人开始报数（从1到3报数），
// 凡报到3的人退出圈子，问最后留下的是原来第几号的那位。（切片实现）
func Question18() {
	var n int = 7
	//fmt.Print("请输入人数:")
	//fmt.Scan(&n)

	sliceA := []int{}
	for i := 1; i <= n; i++ {
		sliceA = append(sliceA, i)
	}
	question18_func(sliceA[0:]...)
}
func question18_func(sliceA ...int) {
	if len(sliceA) >= 3 {
		for i, _ := range sliceA {
			if (i+1)%3 == 0 {
				sliceA = append(sliceA[i+1:], sliceA[0:i]...)
				break
			}
		}
		question18_func(sliceA[0:]...)
	} else {
		fmt.Println(sliceA)
	}
}

// 19、题目：一个三位数哟0-7组成，求所能组成的奇数个数。
func Question19() {
	var count int
	for i := 1; i < 7; i++ {
		for j := 0; j < 7; j++ {
			for k := 0; k < 7; k += 2 {
				if (i*100+j*10+k)%2 != 0 {
					count++
					// fmt.Println(i *100 + j * 10 + k)
				}
			}
		}
	}
	fmt.Println(count)
}

// 20、题目：随机1-100以内的数10个并且没有重复。
func Question20() {
	rand.Seed(time.Now().UnixNano())
	//randSlice := []int{}
	randSlice := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		v := rand.Intn(100) + 1
		for j := 0; j < len(randSlice); j++ {
			if randSlice[j] == v {
				v = rand.Intn(100) + 1
				j = -1
			}
		}
		randSlice = append(randSlice, v)
	}
	fmt.Println(randSlice)

}

// 21、题目：从键盘输入一个字符串，将小写字母全部转换成大写字母。
func Question21() {
	var str string
	fmt.Scan(&str)
	fmt.Println(strings.ToUpper(str))
}

// 22、题目：从键盘输入一个字符数组，统计各个字符出现的次数。

func Question22() {
	var input string
	fmt.Println("请输入一行字符：")
	fmt.Scan(&input)
	//sort.Strings(input)
	array := []byte(input)
	//var result map[byte]int
	result := make(map[string]int, len(array))
	for i := 0; i < len(array); i++ {
		//fmt.Printf("%c\n", array[i])
		//temp := result[fmt.Sprintf("%c", array[i])]
		//temp++
		//result[fmt.Sprintf("%c", array[i])] = temp
		result[fmt.Sprintf("%c", array[i])]++
	}
	fmt.Println(result)
}

//23、按下面的公式求sum的值。
//sum = 1 - 2 + 3 - 4 + 5 - 6 + ⋯⋯ + 99 - 100

func Question23() {
	sum := 0
	str := " = "
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum -= i
			str += fmt.Sprintf("- %d ", i)
		} else {
			sum += i
			str += fmt.Sprintf("+ %d ", i)
		}
	}
	fmt.Println(sum, str)
}

//24、求[m，n]之间既不能被7整除也不能被5整除的整数之和，m和n的值由键盘输入。
//例如，如果m和n的值分别为10和20，则计算结果为：106。
func Question24() {
	for {
		sum, m, n := 0, 0, 0

		fmt.Scan(&m, &n)
		if m > n {
			m, n = n, m
		}
		for i := m; i <= n; i++ {
			if i%7 != 0 && i%5 != 0 {
				sum += i
			}
		}
		fmt.Println(sum)
	}
}

//25、按下面的公式求sum的值。
//sum = m + (m+1) + (m+2) + (m+3) + ⋯⋯ + (n-1) + n
// 例如，如果m和n的值分别为1和100，则计算结果为5050。
func Question25() {
	sum, m, n := 0, 0, 0
	fmt.Scan(&m, &n)
	if m > n {
		m, n = n, m
	}
	for i := m; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//26、不是用系统函数，求字符串的长度并输出。
//例如，当字符串1为"This Is a c Program" 则应输出：Result is: 19
func Question26() {
	// 方法一
	input := bufio.NewReader(os.Stdin)
	str, _ := input.ReadString('\n')
	////方法二 有问题未解决
	//input := bufio.NewScanner(os.Stdin)
	//str := ""
	//for {
	//	if input.Scan() {
	//		str = input.Text()
	//	}
	//}
	sliceA := []byte(str)
	fmt.Println(len(sliceA) - 1)
}

//27、从键盘上输入两个正整数x,y，求它们的最大公约数。
func Question27() {
	x, y := 0, 0
	fmt.Scan(&x, &y)
	if x < y {
		x, y = y, x
	}
	question27_func(x, y)
}
func question27_func(x, y int) {
	if x%y == 0 {
		fmt.Println("最大公约数：", y)
	} else {
		x, y = y, x
		y %= x
		question27_func(x, y)
	}
}

//28、爱因斯坦曾出过这样一道数学题：有一条长阶梯，若每步跨2阶，最后剩下1阶；
// 若每步跨3阶，最后剩下2阶；若每步跨5阶，最后剩下4阶；若每步跨6阶，最后剩下5阶；
// 只有每步跨7阶，最后才正好1阶不剩。编程打印这条阶梯共有多少阶。
func Question28() {
	i := 1
	for {
		if i%2 == 1 && i%3 == 2 && i%5 == 4 && i%7 == 0 {
			break
		}
		i++
	}
	fmt.Println(i)
}

//29、有个切片，找出第二大的数，并且打印出来
//a []int = []int {5,100,32,45,21,67,32,68,41,99,13,71};
func Question29() {
	sliceA := []int{5, 100, 32, 45, 21, 67, 32, 68, 41, 99, 13, 71}
	for i := 0; i < 2; i++ {
		max := i
		for j := i + 1; j < len(sliceA); j++ {
			if sliceA[max] < sliceA[j] {
				max = j
			}
		}
		sliceA[i], sliceA[max] = sliceA[max], sliceA[i]
	}
	fmt.Println(sliceA[1])
}

// 30，双色球
func Question30() {
	var red [6]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(red); i++ {
		v := getRand(1, 33)
		for j := 0; j < i; j++ {
			if v == red[j] {
				v = getRand(1, 33)
				j = -1
			}
		}
		red[i] = v
	}
	fmt.Println(red, "\t", rand.Intn(16)+1)
}
func getRand(begin, end int) int {
	return rand.Intn(end) + begin
}
