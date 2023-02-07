package targetsum

type Memoizer struct {
	Index int
	Sum   int
}

func Solution(nums []int, target int) int {
	return dp(nums, 0, 0, target, make(map[Memoizer]int, 0))
}

func dp(nums []int, startIndex, accumulate, target int, memoizer map[Memoizer]int) int {
	if v, ok := memoizer[Memoizer{Index: startIndex, Sum: accumulate}]; ok {
		return v
	}

	if startIndex == len(nums) {
		if accumulate == target {
			return 1
		}
		return 0
	}

	ways := 0
	ways += dp(nums, startIndex+1, accumulate+nums[startIndex], target, memoizer)
	ways += dp(nums, startIndex+1, accumulate-nums[startIndex], target, memoizer)

	memoizer[Memoizer{Index: startIndex, Sum: accumulate}] = ways
	return ways
}
