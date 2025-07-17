package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/subArrayProduct"
)

func main() {
	nums := []int{1, 2, 3}
	k := 0
	ans := subArrayProduct.NumSubarrayProductLessThanK(nums, k)
	fmt.Println("printing from main", ans)
}
