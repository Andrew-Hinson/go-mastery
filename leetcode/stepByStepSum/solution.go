package stepByStepSum

//Given an array of integers nums, you start with an initial positive value startValue.
//In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
//Return the minimum positive value of startValue such that the step by step sum is never less than 1.

//Input: nums = [-3,2,-3,4,2]
//Output: 5
//Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.

// Hint 1
// Find the minimum prefix sum.

func minStartValue(nums []int) int {
	//psuedo thinking
	//add startvalue to each int in nums to get "step by step sum"
	//return minimum positive value of startval, only when the sum is more than or equal to 1.
	//	what's the minimum value to get a positive step by step sum.
	prefix := []int{nums[0] + 4}
	ans := 0
	for i := 0; i < len(nums); i++ {
		prefix = append(prefix, nums[i]+prefix[len(prefix)-1])

	}
}
