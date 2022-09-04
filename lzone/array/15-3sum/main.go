package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i, _ := range nums {
		t := 0 - nums[i]
		j := i + 1
		k := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			twoSum := nums[j] + nums[k]

			if twoSum == t {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if twoSum < t {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if twoSum > t {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}

			continue
		}
	}

	return res
}

func main() {
	numsList := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{-1, 0, 1},
		[]int{1, -1, -1, 0},
		[]int{-1, 0, 1, 2, -1, -4},
	}

	for _, nums := range numsList {
		fmt.Println(threeSum(nums))
	}
}
