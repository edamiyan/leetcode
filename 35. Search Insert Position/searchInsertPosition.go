package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for nums[l] != target {
		if r-l == 1 {
			if nums[l] < target {
				return l + 1
			} else if nums[l] > target {
				return l
			}
		}

		mid := (l + r) / 2

		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid
		}
	}
	return l
}

func main() {
	arr := []int{1, 3, 5, 6}
	target := 4
	fmt.Println(searchInsert(arr, target))
}
