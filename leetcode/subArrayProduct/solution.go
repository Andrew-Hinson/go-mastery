package subArrayProduct

import "fmt"

// I had to look up the solution. Feel like an idiot after viewing the solution.

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
	if k <= 1 {
		return 0
	}

	product := 1
	count := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for product >= k && left <= right {
			product /= nums[left]
			left++
		}

		count += right - left + 1
		fmt.Println(count)
	}

	return count
}

//		if right == left {
//			if right == len(validNums)-1 {
//				break
//			}
//			right++
//		}
