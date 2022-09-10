package main

import "fmt"

func maxNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func canJump(nums []int) bool {
	len := len(nums)
	max := 0
	for i := 0; i <= max; i++ {
		max = maxNum(max, i+nums[i])
		if max >= len-1 {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 4, 3, 1, 4}
	fmt.Println(canJump(nums))
}
