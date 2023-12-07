package main

import (
	"fmt"
)

var case1 = [][]int{
	{52, 426},
	{94, 1374},
	{75, 1279},
	{94, 1216},
}

func main() {
	result := solve(case1)
	fmt.Println("result:", result)

	part2result := solve([][]int{
		{52947594, 426137412791216},
	})

	fmt.Println("result 2:", part2result)
}

func solve(races [][]int) int {
	result := 1
	for _, rd := range races {
		winway := 0
		first := 0
		for hold := 1; hold <= rd[0]; hold += 1 {
			restTime := rd[0] - hold
			distance := restTime * hold
			if distance >= rd[1] {
				first = hold
				break
			}
		}
		last := 0
		for hold := rd[0] - 1; hold > first; hold -= 1 {
			restTime := rd[0] - hold
			distance := restTime * hold
			if distance >= rd[1] {
				last = hold
				break
			}
		}

		winway = last - first + 1
		result *= winway
	}
	return result
}
