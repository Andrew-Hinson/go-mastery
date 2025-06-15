package sortedArraySquares

import (
	"fmt"
	"sort"
)

//Given an integer array nums sorted in non-decreasing order,
//return an array of the squares of each number sorted in non-decreasing order.
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].

func SortedSquares(nums []int) []int {
	var sortedSlice []int
	for _, num := range nums {
		fmt.Println("iterating on: ", num)
		squaredNum := num * num
		fmt.Println("appending num: ", squaredNum)
		sortedSlice = append(sortedSlice, squaredNum)
	}
	fmt.Println("unsortedNums: ", sortedSlice)
	sort.Ints(sortedSlice)
	fmt.Println("sortedNums: ", sortedSlice)
	return sortedSlice
}
