package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/majority-element/description/
/*Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2*/
func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 2}
	sort.Ints(nums)
	fmt.Println("majorityElementMooresVoting is:", majorityElementMooresVoting([]int{1, 1, 1, 2, 2, 2, 2}))
	fmt.Println("majorityElementOptimised is", majorityElementOptimised(nums))
	fmt.Println("majorityElementBrute is", majorityElementBrute([]int{3, 2, 3}))
}

func majorityElementBrute(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[i]
		}

	}
	return -1

}

func majorityElementOptimised(nums []int) int {

	freq := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			freq++

		} else {
			freq = 1

		}
		ans = nums[i]
	}
	fmt.Println("freq", freq, "ans", ans, len(nums)/2)
	if freq > len(nums)/2 {
		return ans
	}
	return -1
}

func majorityElementMooresVoting(nums []int) int {
	freq := 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if freq == 0 {
			ans = nums[i]
		}
		if ans == nums[i] {
			freq++

		} else {
			freq--

		}

	}
	return ans
}
