package combination_sum

import (
	"sort"
)

func Solution(candidates []int, target int) (ans [][]int) {
	if len(candidates) == 0 {
		return
	}
	sort.Ints(candidates)
	recur(&ans, make([]int, 0), candidates, target, 0)
	return
}

func recur(ans *[][]int, combinations, candidates []int, target, startIndex int) {
	if target < 0 || startIndex > len(candidates) {
		return
	}

	if target == 0 {
		cpyCombinations := make([]int, len(combinations))
		copy(cpyCombinations, combinations)
		*ans = append(*ans, cpyCombinations)
	}

	for i := startIndex; i < len(candidates); i++ {
		combinations = append(combinations, candidates[i])
		recur(ans, combinations, candidates, target-candidates[i], i)
		combinations = combinations[:len(combinations)-1]
	}
}
