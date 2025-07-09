package kRadiusAverages

import "fmt"

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

func GetAverages(nums []int, k int) []int {
	//holds the subarray total
	totals := []int{nums[0]}
	averages := []int{}
	sum := 0
	divisor := k*2 + 1

	for i := 1; i < len(nums); i++ {
		//nums[i]+averages[len(averages)-1] = last element of averages
		totals = append(totals, (nums[i])+totals[len(totals)-1])
	}
	fmt.Println(totals)

	//I think I iterate over totals and compute the averages with k radius
	for j := 0; j < len(nums); j++ {
		//if j is less than k, j equals -1.
		//if the len(total) - k, j equals -1
		if j < k {
			averages = append(averages, -1)
			continue
		}
		if j >= k && j < len(totals)-k {
			//	sum exists in totals, to get total to average, get last element +k - nums[k*2]
			//The above thought is incorrect, its k * 2 + 1 because it's a radius which assumes equal sides centered on an int
			//my mistake is not calculating the new subarray value, because this current iteration is taking in all of the numbers and not subtracting numbers that aren't included
			//I think I need to calc each subararray separately
			//I dont need to calc separately just handle the totals[0] situation
			if totals[j-k] == totals[0] {
				sum = totals[j+k]
				averages = append(averages, sum/divisor)
				continue
			}
			sum = totals[j+k] - totals[j-k]
			averages = append(averages, sum/divisor)

			continue
		}
		averages = append(averages, -1)

	}

	return averages
}
