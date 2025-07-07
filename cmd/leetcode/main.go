package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/stepByStepSum"
)

func main() {
	nums := []int{-3, 2, -3, 4, 2}

	ans := stepByStepSum.MinStepByStep(nums)
	fmt.Println("Answer value: ", ans)
}
