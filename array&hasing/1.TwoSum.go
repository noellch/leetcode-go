/* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order. */

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	table := map[int]int{}

	for i, num := range nums {
		if val, found := table[num]; found {
			return []int{val, i}
		}

		table[target-num] = i
	}

	return nil
}
