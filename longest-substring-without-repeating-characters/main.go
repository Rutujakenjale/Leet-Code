package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstringOptimized(s))

}

func lengthOfLongestSubstringOptimized(s string) int {
	m := make(map[byte]int, len(s))
	right, left, maxLen := 0, 0, 0

	for right < len(s) {
		if val, ok := m[s[right]]; ok {
			left = max(left, val+1)
		}
		m[s[right]] = right

		maxLen = max(maxLen, right-left+1)
		right++

	}
	return maxLen

}
func lengthOfLongestSubstringBrute(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]bool, len(s))
		for j := i; j < len(s); j++ {
			if m[s[j]] {
				break
			}
			m[s[j]] = true
			if (j-i)+1 > maxLen {
				maxLen = j - i + 1
			}

		}
	}
	return maxLen
}
