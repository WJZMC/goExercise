package algorithms

import "fmt"

//稀疏数组
type node struct {
	row   int
	col   int
	value int
}

func sparseArray() {
	var oriArray [11][11]int
	oriArray[1][1] = 2
	oriArray[2][2] = 3

	//稀疏数组记录
	var sparseA []node
	sparseA = append(sparseA, node{11, 11, 0})
	for i, v := range oriArray {
		for j, v2 := range v {
			fmt.Printf("%d ", v2)
			if v2 != 0 {
				tmp := node{i, j, v2}
				sparseA = append(sparseA, tmp)
			}
		}
		fmt.Println()
	}

	fmt.Println("稀疏数组存储：", sparseA)
	//稀疏数组还原

	fmt.Println("稀疏数组还原：")

	var row = sparseA[0].row
	var col = sparseA[0].col
	var value = sparseA[0].value
	var reset [][]int
	for i := 0; i < row; i++ {
		var temp []int
		for j := 0; j < col; j++ {
			temp = append(temp, value)
		}
		reset = append(reset, temp)
	}

	for i, v := range sparseA {
		if i > 0 {
			reset[v.row][v.col] = v.value
		}
	}

	for _, v := range oriArray {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

}
