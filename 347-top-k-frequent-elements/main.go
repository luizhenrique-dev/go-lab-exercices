package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 2, 5, 5, 5, 6}
	k := 2

	result := topKFrequent(nums, k)
	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {
	numsByOccurrences := buildMapNumbersByOccurrences(nums)
	buckets := buildBuckets(nums, numsByOccurrences)
	return getTopFrequent(k, buckets)
}

func getTopFrequent(k int, buckets [][]int) []int {
	toReturn := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(toReturn) < k; i-- {
		if buckets[i] != nil {
			toReturn = append(toReturn, buckets[i]...)
		}
	}
	return toReturn
}

func buildMapNumbersByOccurrences(nums []int) map[int]int {
	numsByOccurrences := make(map[int]int)
	for _, currentNum := range nums {
		numsByOccurrences[currentNum] += 1
	}
	return numsByOccurrences
}

func buildBuckets(nums []int, numsByOccurrences map[int]int) [][]int {
	buckets := make([][]int, len(nums)+1)
	for number, count := range numsByOccurrences {
		if buckets[count] == nil {
			buckets[count] = make([]int, 0)
		}
		buckets[count] = append(buckets[count], number)
	}
	return buckets
}

func topKFrequent2(nums []int, k int) []int {
	numsByOccurrences := make(map[int]int)
	for _, currentNum := range nums {
		numsByOccurrences[currentNum] += 1
	}

	toReturn := make([]int, 0, k)
	for i := 0; i < k; i++ {
		var topFrequentNumber, topOccurrences int
		for number, count := range numsByOccurrences {
			if count > topOccurrences {
				topOccurrences = count
				topFrequentNumber = number
			}
		}
		toReturn = append(toReturn, topFrequentNumber)
		delete(numsByOccurrences, topFrequentNumber)
	}
	return toReturn
}
