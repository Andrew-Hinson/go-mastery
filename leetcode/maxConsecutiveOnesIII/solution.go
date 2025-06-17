package maxConsecutiveOnesIII

//Given a binary array nums and an integer k,
//return the maximum number of consecutive 1's in the array
//if you can flip at most k 0's.
//Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
//Output: 6
//Explanation: [1,1,1,0,0,1,1,1,1,1,1]
//Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// FindMaxConsecutiveOnes = basically, what is the greatest number between right and left
func FindMaxConsecutiveOnes(ones []int, k int) int {
	left := 0
	ans := 0
	//current is the number of 0s in the window
	curr := 0
	for right := range ones {
		if ones[right] == 0 {
			curr += 1
		}
		//while right has added more than 2 zeros to curr, bring left along with it till it gets to 0s and subtract from curr
		for curr > k {
			if ones[left] == 0 {
				curr -= 1
			}
			//fixed mistake.
			left += 1
		}
		//another gotcha, this is cleaner. and skips the self assignment issue.
		if ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return ans
}
