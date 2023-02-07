package targetsum

func Solution(nums []int, target int) int {
	return dp(nums, 0, 0, target)
}

func dp(nums []int, startIndex, accumulate, target int) int {
	if startIndex == len(nums) {
		if accumulate == target {
			return 1
		}
		return 0
	}

	ways := 0
	ways += dp(nums, startIndex+1, accumulate+nums[startIndex], target)
	ways += dp(nums, startIndex+1, accumulate-nums[startIndex], target)
	return ways
}
