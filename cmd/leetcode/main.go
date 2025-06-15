package main

import (
	"github.com/Andrew-Hinson/go-mastery/leetcode/sortedArraySquares"
)

func main() {
	initialSlice := []int{-4, -1, 0, 3, 10}
	print(initialSlice)
	result := sortedArraySquares.SortedSquares(initialSlice)
	print(result)
}
