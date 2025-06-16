package main

import maxsubarray1 "github.com/Andrew-Hinson/go-mastery/leetcode/maxSubArray-1"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	ans := maxsubarray1.FindMaxAverage(nums, k)
	println(ans)
}
