package main

import "fmt"

/*
nums1 = [1,2,3,0,0,0]
m = 3
nums2 = [2,5,6]
n = 3
response = merge_lists(nums1, m, nums2, n)
print(response)

# Arrays de prueba 2
nums1 = [1,2,3,0,0,0,0]
m = 3
nums2 = [-4,2,3,9]
n = 4

- Set a pointer for the ordered list
- Set a pointer for the left list
- Set a pointer for the right list
- Loop until both pointer reach 0
*/

func mergeLists(numsLeft, numsRight []int, n, m int) []int {
	cP := n + m - 1
	lP := n - 1
	rP := m - 1

	for lP >= 0 && rP >= 0 {
		if numsLeft[lP] > numsRight[rP] {
			numsLeft[cP] = numsLeft[lP]
			lP--
			cP--
		} else {
			numsLeft[cP] = numsRight[rP]
			rP--
			cP--
		}
	}

	if n != m {
		for i, v := range numsRight[:rP+1] {
			numsLeft[i] = v
		}
	}
	return numsLeft
}

func main() {
	numsOne := []int{1, 2, 3, 0, 0, 0, 0}
	numsTwo := []int{-4, 2, 3, 9}
	n := 3
	m := 4
	numsOne = mergeLists(numsOne, numsTwo, n, m)
	fmt.Println(numsOne)

}
