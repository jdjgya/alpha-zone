package main

import "fmt"

func minIndex(m map[byte]int) int {
	count := 0
	min := 0

	for _, v := range m {
		if count == 0 {
			min = v
			count++
			continue
		}

		if v < min {
			min = v
		}
	}

	return min
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstringTwoDistinct(str string) int {
	maxLen := 0
	charMap := map[byte]int{}

	if len(str) <= 1 {
		return 1
	}

	for left, right := 0, 0; right < len(str); right++ {
		c := str[right]
		charMap[c] = right

		distinctChar := len(charMap)
		if distinctChar > 2 {
			smallIndex := minIndex(charMap)
			delete(charMap, str[smallIndex])
			left = smallIndex + 1
			continue
		}

		maxLen = maxNum(right-left+1, maxLen)
	}

	if len(charMap) == 1 && len(str) > 1 {
		return len(str)
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ee"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("e"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("cdaba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abaccc"))
}
