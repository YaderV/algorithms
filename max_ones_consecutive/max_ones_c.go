package main

import (
	"fmt"
	"math"
)

/*
	- Have two pointers start, end = 0
	- Move end to the right
	- Count the amount of 1 (ones)
	- Calc the max length (start - end + 1)
	- Calc the amount of flipped values (length - ones)
	- When the amount of fillped values > k (it means we have flipped k times some 0s)
	start a new loop that move the start pointer foward
	- If the new start value is 1, decrease number of 1s.
	- The inner loop ends, when we are able to flip 0 again; (length - ones > k); then continue with the
	max length calc and moving the right pointer
*/

func maxOneConsecutive(nums []int, k int) int {
	var start, end, max, ones int

	for end <= len(nums)-1 {
		if nums[end] == 1 {
			ones++
		}

		// length - ones = number of flipped 0
		// we continue with the calc until we find a range where
		// we are allow to flip again
		for (end-start+1)-ones > k {
			// We are going to move the left pointer to calc
			// new possible consecutives 1, so we found a one
			// we remove it from the ones liest
			if nums[start] == 1 {
				ones--
			}
			// move the left pointer
			start++
		}
		max = int(math.Max(float64(max), float64(end-start+1)))
		end++
	}
	return max
}

func main() {
	// k := 3
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(maxOneConsecutive(nums, k))
}
