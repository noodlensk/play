package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var max, j int
	var runes = [128]int{}
	for i := 0; j < len(s); j++ {
		r := int(s[j])
		if d := runes[r]; d > 0 {
			if i < d {
				i = d
			}
		}
		if cur := j - i + 1; max < cur {
			max = cur
		}
		runes[r] = j + 1
	}
	return max
}
func main() {
	fmt.Printf("%+v", lengthOfLongestSubstring("anviaj"))
}
