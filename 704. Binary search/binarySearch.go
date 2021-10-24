package main

import "fmt"

func search(nums []int, target int) int {
	left, right := -1, len(nums)-1

	for left < right-1 {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[right] != target {
		return -1
	}

	return right
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Print(search(arr, 3))
}
