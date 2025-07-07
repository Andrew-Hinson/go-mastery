package stepByStepSum

//Given an array of integers nums, you start with an initial positive value startValue.
//In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
//Return the minimum positive value of startValue such that the step by step sum is never less than 1.

//Input: nums = [-3,2,-3,4,2]
//Output: 5
//Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.

// Hint 1
// Find the minimum prefix sum.

func MinStepByStep(nums []int) int {
	//psuedo thinking
	//add startvalue to each int in nums to get "step by step sum"
	//return minimum positive value of startval, only when the sum is more than or equal to 1.
	//	what's the minimum value to get a positive step by step sum.

	//Newest thought - the MINIMUM STARTVALUE. I.E. What we started the prefix on.
	//reworded problem in my own words:
	//create a prefix array of the nums, one of the values of nums will be our starting number to build the prefix array
	// return the nums[index] that gives us a prefix array where the numbers in the prefix array are never less than 1
	//the index is used as the starting value - sort of: [1, 2, 3, 4, 5] = [-3, 2, -3, 4, 2] (non 0 index I guess)

	//how to solve, iterate through nums, build sum of prefix array
	//if the value of the index is greater than or equal to 1, calculate a prefix array
	//then determine if the values in the prefix array are greater than or equal to 1

	//!----- I admitted defeat after an hour and looked up how to solve this -----!
	// the problem SHOULD be worded (found online):
	//"What's the minimum number I need to add to my worst-case scenario to keep the running sum at least 1?"
	sum := 0
	startValue := 0
	for _, value := range nums {
		//instead of creating an actual slice like previous problems
		// the value is added to the sum
		sum += value
		//if the startValue is greater less than the sum, keep the sum, else keep the startValue
		startValue = myMin(startValue, sum)
	}
	//given the corrected problem statement I added at the top, we flip the worst case negative value of what we are given
	//to a positive number, and add 1. That gives us the starting value needed in order for our running sum to at least be 1.
	//I understand this now, and I still hate it.
	return -startValue + 1
}

// this is a simple minimum value calculation
func myMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
