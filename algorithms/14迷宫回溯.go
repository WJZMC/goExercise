package algorithms

import "fmt"

const MapSize = 8

func maze() {
	//假设1：障碍物    2：已走   0：可走
	mazeMap := [MapSize][MapSize]int{}
	for i := 0; i < MapSize; i++ {
		mazeMap[0][i] = 1
		mazeMap[MapSize-1][i] = 1
		mazeMap[i][0] = 1
		mazeMap[i][MapSize-1] = 1
	}

	//设置障碍物
	mazeMap[3][1] = 1
	mazeMap[4][2] = 1
	mazeMap[4][3] = 1

	listMap(&mazeMap)
	fmt.Println("begin----------")

	mazeTest(&mazeMap, 1, 1)
	fmt.Println("end----------")

	listMap(&mazeMap)

}

func listMap(maze *[MapSize][MapSize]int) {
	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			fmt.Printf("%v", maze[i][j])
		}
		fmt.Println()
	}
}

//假设从【1，1】-->[6,6]
//优先级：下右上左
//
func mazeTest(maze *[MapSize][MapSize]int, i, j int) bool {
	if maze[6][6] == 2 {
		return true
	}
	if maze[i][j] == 0 {
		maze[i][j] = 2              //如果可以走，则置为2
		if mazeTest(maze, i+1, j) { //尝试着往下走
			return true
		} else if mazeTest(maze, i, j+1) {
			return true
		} else if mazeTest(maze, i-1, j) {
			return true
		} else if mazeTest(maze, i, j-1) {
			return true
		}
		maze[i][j] = 3
		return false
	} else {
		return false
	}

}
