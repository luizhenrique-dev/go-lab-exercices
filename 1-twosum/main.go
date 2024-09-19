package main

import "fmt"

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result[0], result[1])

	result = twoSumAlternative(nums, target)
	fmt.Println(result[0], result[1])
}

// twoSum function
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func twoSumAlternative(nums []int, target int) []int {
	indices := make([]int, 2)
	var found bool
	for i, num := range nums {
		if found {
			break
		}
		for j, numSec := range nums {
			if i == j {
				continue
			}
			if num+numSec == target {
				indices[0] = i
				indices[1] = j
				found = true
				break
			}
		}
	}
	return indices
}
