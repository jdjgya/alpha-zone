package main

import "fmt"

func swap(a, b int, nums []int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}

func reverse(left, right int, nums []int) {
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()

	touchDesced := false
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 && !touchDesced {
			reverse(i, len(nums)-1, nums)
			return
		}

		if nums[i-1] >= nums[i] {
			continue
		} else {
			touchDesced = true
		}

		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] > nums[i-1] {
				swap(j, i-1, nums)
				break
			}
		}

		reverse(i, len(nums)-1, nums)
		return
	}
}

func main() {
	nums1 := []int{1, 2, 4, 6, 5, 3}
	nextPermutation(nums1)

	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)

	nums3 := []int{5, 1, 1}
	nextPermutation(nums3)
}
