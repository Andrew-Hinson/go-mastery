package subArrayProduct

import "fmt"

//Given an array of integers nums and an integer k,
//return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.

//Example 1:
//Input: nums = [10,5,2,6], k = 100
//Output: 8
//Explanation: The 8 subarrays that have product less than 100 are:
//[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
//Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

//Example 2:
//Input: nums = [1,2,3], k = 0
//Output: 0

//Input: nums = [10,5,2,6], k = 100
//Output: 8

func NumSubarrayProductLessThanK(nums []int, k int) int {
	//I need to build a product array maybe?
	//sliding window problem methinks
	// I need to iterate through nums,
	// If the number is < k, add it to the product array
	//multiply nums together, for each successful multiply that's less than 100, I increment a counter.
	n := len(nums)
	validNums := make([]int, n+1)
	left := 0
	right := 0
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			counter++
			validNums[i] = nums[i]
		}
	}
	for j := 0; j < len(validNums); j++ {
		if validNums[left]*validNums[right] < k {
			right++
			counter++
			fmt.Printf("{%v,%v}", validNums[left], validNums[right])
			continue
		}
		left++
	}

	return counter
}
