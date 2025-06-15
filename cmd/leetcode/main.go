package main

import (
	"fmt"
	"github.com/Andrew-Hinson/go-mastery/leetcode/sortedArraySquares"
)

func main() {
	initialSlice := []int{-4, -1, 0, 3, 10}
	result := sortedArraySquares.SortedSquares(initialSlice)
	fmt.Println(result)
}
