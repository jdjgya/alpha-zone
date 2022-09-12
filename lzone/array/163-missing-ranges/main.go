package main

import (
	"fmt"
	"strconv"
)

func transRange(lower, upper int) string {
	if lower == upper {
		return strconv.Itoa(lower)
	}

	return strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	result := []string{}
	prev := lower - 1

	for i := 0; i < len(nums)+1; i++ {
		curr := 0
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}

		if prev+1 <= curr-1 {
			result = append(result, transRange(prev+1, curr-1))
		}

		prev = curr
	}

	return result
}

func main() {
	nums := []int{0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, 0, 99))

	nums = []int{-1}
	fmt.Println(findMissingRanges(nums, -1, -1))

	nums = []int{}
	fmt.Println(findMissingRanges(nums, 1, 1))

	nums = []int{}
	fmt.Println(findMissingRanges(nums, -3, 1))
}
