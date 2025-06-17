package maxSubArray_1

import "math"

//You are given an integer array nums consisting of n elements, and an integer k.
//Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
//Any answer with a calculation error less than 10-5 will be accepted.
//Example 1:
//Input: nums = [1,12,-5,-6,50,3], k = 4
//Output: 12.75000
//Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
//Example 2:
//Input: nums = [5], k = 1
//Output: 5.00000

func FindMaxAverage(nums []int, k int) float64 {
	//I initialized curr to be 0, it will hold the initial value of the first k sub array added together
	var curr = 0
	//I iterate through the range of k in nums and add them together and store in curr to get our starting point
	for index := range k {
		curr += nums[index]
	}
	println(curr)
	//ans will be our starting highest average value
	ans := float64(curr) / float64(k)
	println("answer: ", ans)
	for j := k; j <= len(nums)-1; j++ {
		curr += nums[j] - nums[j-k]
		println("current after adding: ", curr)
		newCurrent := float64(curr) / float64(k)
		println("NewCurrent: ", newCurrent)
		ans = math.Max(ans, newCurrent)
	}

	println("answer is now: ", ans*1000000000)

	return ans
}

//Final thoughts, I didn't know how to use math.Max, so while using, realized it required float64s,
//I realized I needed to return a float64, so then just converted the items I would send into it into float64
