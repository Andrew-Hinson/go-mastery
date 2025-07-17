package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/subArrayProduct"
)

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	ans := subArrayProduct.NumSubarrayProductLessThanK(nums, k)
	fmt.Println("printing from main", ans)
}
