package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/runningSum"
)

func main() {
	nums := []int{1, 2, 3, 4}

	ans := runningSum.FindRunningSumOf1dArray(nums)
	fmt.Println(ans)
}
