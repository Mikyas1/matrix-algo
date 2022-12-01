package main

import "fmt"

type Matrix [4][4]int

func (m *Matrix) String() string {
	res := ""
	for _, row := range m {
		for _, col := range row {
			res += fmt.Sprintf("%v ", col)
		}
		res += "\n"
	}
	return res
}

var matrix = Matrix{
	{1, 1, 0, 1},
	{1, 0, 1, 0},
	{1, 1, 1, 0},
	{0, 0, 1, 1},
}

var path = Matrix{}

func findPath(i, j int) int {
	if i == len(matrix)-1 && j == len(matrix[0])-1 {
		path[i][j] = 1
		return 1
	}
	if matrix[i][j] == 1 {
		path[i][j] = 1

		// go right
		if findPath(i, j+1) == 1 {
			return 1
		}

		// go down
		if findPath(i+1, j) == 1 {
			return 1
		}
		path[i][j] = 0

	}
	return 0
}

func main() {
	findPath(0, 0)
	fmt.Println(path.String())
}
