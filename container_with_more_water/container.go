package main

import (
	"fmt"
	"math"
)

func maxArea(h []int) int {
	pLeft := 0
	pRight := len(h) - 1
	max := 0

	for pLeft < pRight {
		area := (pRight - pLeft) * int(math.Min(float64(h[pLeft]), float64(h[pRight])))
		if area > max {
			max = area
		}

		if h[pLeft] > h[pRight] {
			pRight--
		} else {
			pLeft++
		}
	}

	return max

}

func main() {
	// h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	h := []int{8, 1, 6, 2, 5, 4, 1, 3, 7}
	fmt.Println(maxArea(h))
}
