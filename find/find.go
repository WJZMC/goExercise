package find

import "fmt"

func Find() {

	arr := [10]int{5, 10, 20, 50, 100, 500, 600, 700, 1000, 1001}
	BinaryFind(arr, 500, 0, len(arr)-1)
}

func BinaryFind(arr [10]int, findObj, leftIndex, rightIndex int) {
	if leftIndex > rightIndex {
		fmt.Println("没有该数")
	} else {
		middle := (leftIndex + rightIndex) / 2
		if findObj == arr[middle] {
			fmt.Println("找到了下标为：", middle, "的就是要找的数", findObj)
		} else {
			if findObj > arr[middle] {

				BinaryFind(arr, findObj, middle+1, rightIndex)
			} else {
				BinaryFind(arr, findObj, leftIndex, middle-1)
			}
		}
	}

}
