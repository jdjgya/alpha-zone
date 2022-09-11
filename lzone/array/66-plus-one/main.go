package main

import "fmt"

func plusOne(digits []int) []int {
	end := len(digits) - 1

	for i := end; i >= 0; i-- {
		digits[i]++
		if digits[i]/10 == 0 {
			return digits
		}

		if i == 0 {
			digits = append([]int{0}, digits...)
			i++
		}

		digits[i] = digits[i] % 10
		digits[i-1] = digits[i-1] + (digits[i] / 10)
	}

	return digits
}

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println(plusOne(nums))

	nums = []int{3, 9, 9, 9}
	fmt.Println(plusOne(nums))

	nums = []int{9}
	fmt.Println(plusOne(nums))

	nums = []int{9, 9}
	fmt.Println(plusOne(nums))

	nums = []int{0}
	fmt.Println(plusOne(nums))
}
