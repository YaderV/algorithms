package main

import "fmt"

func swipe(data []int, from, to int) {
	tmp := data[to]
	data[to] = data[from]
	data[from] = tmp
}

// [0, 0, 2, 1, 1, 2]
//

func sortColor(colors []int) {
	index := 0
	leftP := 0
	rightP := len(colors) - 1

	for index < rightP {
		if colors[index] == 0 {
			swipe(colors, index, leftP)
			leftP++
		} else if colors[index] == 2 {
			swipe(colors, index, rightP)
			rightP--
			index--
		}
		index++
	}
}

func main() {
	list := []int{2, 0, 2, 1, 1, 0}
	sortColor(list)
	fmt.Println(list)
}
