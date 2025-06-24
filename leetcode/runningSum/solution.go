package runningSum

//Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//
//Return the running sum of nums.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,4]
//Output: [1,3,6,10]
//Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

func FindRunningSumOf1dArray(nums []int) []int {
	prefix := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		println("I = ", i)
		prefix = append(prefix, nums[i]+len(prefix)-1)
	}
	return nums
}
