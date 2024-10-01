package main

import "fmt"

func main() {
	// nums := []int{100, 4, 200, 1, 3, 2}
	// nums := []int{100, 4, 200, 1, 2, -5, 3, -1, -2, -3, -4, 0}
	// nums := []int{0, 0}
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	l := longestConsecutive(nums)
	fmt.Println(l)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for findNextNumber(currentNum, numSet) {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numbersSetMap := make(map[int]bool)
	currentStreakMap := make(map[int]bool)
	longestStreakMap := make(map[int]bool)
	var currentNum int

	for _, num := range nums {
		numbersSetMap[num] = true
	}
	for _, num := range nums {
		if longestStreakMap[num] {
			continue
		}
		if numbersSetMap[num+1] {
			currentNum = num
			currentStreakMap[currentNum] = true
			for findNextNumber(currentNum, numbersSetMap) {
				currentNum += 1
				currentStreakMap[currentNum] = true
			}

			if len(currentStreakMap) > len(longestStreakMap) {
				longestStreakMap = currentStreakMap
			}
			currentStreakMap = make(map[int]bool)
		}

	}

	if len(longestStreakMap) == 0 {
		return 1
	}

	return len(longestStreakMap)
}

func findNextNumber(currentNum int, numberSetMap map[int]bool) bool {
	return numberSetMap[currentNum+1]
}
