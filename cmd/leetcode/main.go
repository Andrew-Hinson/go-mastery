package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/kRadiusAverages"
)

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	ans := kRadiusAverages.GetAverages(nums, k)
	fmt.Println("printing from main", ans)
}
