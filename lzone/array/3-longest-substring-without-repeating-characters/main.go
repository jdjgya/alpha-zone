package main

import (
	"fmt"
)

func lengthOfLongestSubstring(str string) int {
	maxLen := 0
	location := make(map[byte]int)

	for start, end := 0, 0; end < len(str); end++ {
		char := str[end]
		if index, found := location[char]; found && index < end {
			if index+1 > start {
				start = index + 1
			}
		}

		location[char] = end

		if (end - start + 1) > maxLen {
			maxLen = (end - start + 1)
		}
	}

	return maxLen
}

func main() {
	str1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str1))

	str2 := "bbbbbbbb"
	fmt.Println(lengthOfLongestSubstring(str2))

	str3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str3))

	str4 := "dvdf"
	fmt.Println(lengthOfLongestSubstring(str4))

	str5 := "ohomm"
	fmt.Println(lengthOfLongestSubstring(str5))

	str6 := "nfpdmpi"
	fmt.Println(lengthOfLongestSubstring(str6))

	str7 := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(str7))

	str8 := "abba"
	fmt.Println(lengthOfLongestSubstring(str8))
}
