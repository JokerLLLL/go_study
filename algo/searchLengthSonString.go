package algo

import "fmt"

func SearchLengthSonString(s string) int {
	now := 0
	max := 0
	mapLocationLast := make(map[rune]int)
	ru := []rune(s)
	for i, v := range ru {
		if i2, ok := mapLocationLast[v]; ok && i2 >= now {
			now = i2 + 1
		} else {
			if i - now + 1 > max {
				max = i - now + 1
			}
		}
		mapLocationLast[v] = i
	}
	fmt.Println(max)
	return max
}