package main

import (
	"github.com/Andrew-Hinson/go-mastery/leetcode/maxConsecutiveOnesIII"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	ans := maxConsecutiveOnesIII.FindMaxConsecutiveOnes(nums, k)
	println(ans)
}
