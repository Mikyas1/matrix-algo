package main

import "fmt"

type Matrix [5][5]int

func (m Matrix) String() string {
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
	{1, 0, 0, 1, 0},
	{1, 0, 1, 0, 0},
	{0, 0, 1, 0, 1},
	{1, 0, 1, 0, 1},
	{1, 0, 1, 1, 0},
}

// answer [1, 2, 2, 2, 5]

func riversLength(matrix Matrix) []int {
	var ans []int
	var seen Matrix
	for r, row := range matrix {
		for c, _ := range row {
			var count int
			traverseThroughRiver(r, c, &matrix, &seen, &count)
			if count > 0 {
				ans = append(ans, count)
			}
		}
	}
	//fmt.Println(fmt.Sprintf("%s", seen))
	return ans
}

func traverseThroughRiver(r int, c int, matrix, seen *Matrix, count *int) {
	if seen[r][c] == 1 {
		return
	}
	seen[r][c] = 1
	if matrix[r][c] == 0 {
		return
	}

	*count++

	// do depth first search
	// get neighbours
	neighbors := getNeighbours(r, c, matrix)
	for _, neighbors := range neighbors {
		traverseThroughRiver(neighbors[0], neighbors[1], matrix, seen, count)
	}
}

func getNeighbours(r int, c int, matrix *Matrix) [][]int {
	var neighbours [][]int
	if r != 0 {
		neighbours = append(neighbours, []int{r - 1, c})
	}
	if r != len(matrix)-1 {
		neighbours = append(neighbours, []int{r + 1, c})
	}
	if c != 0 {
		neighbours = append(neighbours, []int{r, c - 1})
	}
	if c != len(matrix)-1 {
		neighbours = append(neighbours, []int{r, c + 1})
	}
	return neighbours
}

func main() {
	ans := riversLength(matrix)
	fmt.Println(ans)
}
