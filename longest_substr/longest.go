package main

import "fmt"

func longestSubStr(str string) int {
	max := 0
	repeated := make(map[string]bool)
	lPointer := 0
	rPointer := 0

	for rPointer < len(str)-1 {
		letter := string(str[rPointer])
		isRepeated, ok := repeated[letter]
		// If we don't find the lletter
		if !isRepeated || !ok {
			repeated[letter] = true
			rPointer++
		} else {
			substrLenght := rPointer - lPointer
			if substrLenght > max {
				max = substrLenght
			}
			lPointer++
			rPointer = lPointer
			repeated = make(map[string]bool)
		}
	}

	if rPointer-lPointer > max {
		max = rPointer - lPointer
	}
	return max
}

func main() {
	// str := "abcabcbb"
	str := "jdkafnlcdsalkxcmpoiuytfccv"
	fmt.Println(longestSubStr(str))
}
