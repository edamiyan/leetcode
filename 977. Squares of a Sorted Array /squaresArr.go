package main

import "fmt"

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return n * -1
}

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	leftIndex := 0
	rightIndex := len(nums) - 1
	resultIndex := rightIndex
	for rightIndex-leftIndex > -1 {
		if Abs(nums[leftIndex]) > Abs(nums[rightIndex]) {
			result[resultIndex] = nums[leftIndex] * nums[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = nums[rightIndex] * nums[rightIndex]
			rightIndex--
		}
		resultIndex--
	}
	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Print(sortedSquares(nums))
}
