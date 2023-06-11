package main

import (
	"fmt"
	"math"
)

// create a list with the max left height for each elements
// create a list with the max right height for each elements
// for each element cal the amount of water (min(right) - min(left)) - current height
// sum each amount of water

func trapping(h []int) int {
	leftH := make([]int, len(h))
	rightH := make([]int, len(h))
	var max int
	for i := 0; i < len(h); i++ {
		if h[i] > max {
			max = h[i]
		}
		leftH[i] = max
	}
	max = 0
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] > max {
			max = h[i]
		}
		rightH[i] = max
	}

	var water int
	for i, v := range h {
		water += int(math.Min(float64(leftH[i]), float64(rightH[i]))) - v
	}
	return water
}

func main() {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapping(heights))

}
