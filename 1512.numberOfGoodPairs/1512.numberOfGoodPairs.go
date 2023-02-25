package main

import (
	"fmt"
	"leet-by-bit/util/msg"
)

func numIdenticalPairs(nums []int) int {
	count := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		msg.Info(">>>>>>>>>>>>>>>>>>>>")
		msg.Info("i", i, " j", j)
		msg.Info("====================")
		for j < len(nums) {
			msg.Info("i", i, " j", j)
			if i < j {
				if nums[i] == nums[j] {
					msg.Info("nums i", nums[i], " nums j", nums[j])
					count++
				}
			}
			j++
		}
		j = 0
	}
	return count
}

func main() {
	msg.HideInfo(true)

	fmt.Println("result ", numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println("result ", numIdenticalPairs([]int{1, 1, 1, 1}))       // 6
	fmt.Println("result ", numIdenticalPairs([]int{1, 2, 3}))          // 0
}

// fastest solution
