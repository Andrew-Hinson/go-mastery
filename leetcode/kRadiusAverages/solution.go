package kRadiusAverages

// You are given a 0-indexed array nums of n integers, and an integer k.
// The k-radius average for a subarray of nums centered at some index i
// with the radius k is the average of all elements in nums between the indices i - k
// and i + k (inclusive). If there are less than k elements before or after the index i,
// then the k-radius average is -1.
// Build and return an array avgs of length n where avgs[i] is the k-radius average for the
// subarray centered at index i.
// The average of x elements is the sum of the x elements divided by x, using integer division.
// The integer division truncates toward zero, which means losing its fractional part.
// For example, the average of four elements 2, 3, 1, and 5 is (2 + 3 + 1 + 5) / 4 = 11 / 4 = 2.75,
// which truncates to 2.
// Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
// Output: [-1,-1,-1,5,4,4,-1,-1,-1]
//        [7 11 14 23 24 32 37 39 45]

//new approach - I need to calculate the subaraary only if the item is a valid number to calculate an average for.
//no i don't I just need to subtract the subarray from the total.

// [18334,25764,19780,92480,69842,73255,89893] given
// [18334,25764,19780,92480,69842,73255,89893] expected
// k = 0

func GetAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	n := len(nums)
	//holds the subarray total
	//key to note here we are adding an element to handle several things
	//if we only allocated n elements the loop for calculating the totals would break
	//the range sum of totals[right+1] - totals[left] requires it for the last element
	totals := make([]int, n+1)
	//will return averages
	averages := make([]int, n)
	// divisor is radius * 2, then adding 1 for itself
	divisor := k*2 + 1
	//instead of calculating what will be -1, just put them all as -1 to start
	for i := range averages {
		averages[i] = -1
	}
	//I was calculating the prefix total incorrectly before
	//I need to do i + 1 to get the total of the 0 index to i in the original array
	for i := 0; i < len(nums); i++ {
		totals[i+1] = totals[i] + nums[i]
	}

	for j := k; j < n-k; j++ {
		left := j - k
		right := j + k
		sum := totals[right+1] - totals[left]
		average := sum / divisor
		averages[j] = average
	}

	return averages
}
