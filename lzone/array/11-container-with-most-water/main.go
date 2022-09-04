package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxZone := 0

	for l < r {
		availH := min(height[r], height[l])
		availW := r - l
		maxZone = max(maxZone, availH*availW)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxZone
}

func main() {
	max := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(max)
}
