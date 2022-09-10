package main

import "fmt"

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
}

func reflect(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

func rotate2(matrix [][]int) {
	transpose(matrix)
	reflect(matrix)
}

func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

func main() {
	mtx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate1(mtx)
	fmt.Println(mtx)

	mtx = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate2(mtx)
	fmt.Println(mtx)
}
