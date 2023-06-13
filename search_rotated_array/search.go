package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			// Check if it is not rotated
			if nums[mid] < nums[end] {
				end = mid - 1
			} else {
				start = mid + 1
			}

		} else {
			// Check if it is not rotated
			if nums[start] < nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
