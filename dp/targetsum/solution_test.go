package targetsum

import (
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		actual int
	}{
		{nums: []int{1, 1, 1, 1, 1}, target: 3, actual: 5},
		{nums: []int{1}, target: 1, actual: 1},
	}

	for _, tc := range testCases {
		name := make([]string, 0, len(tc.nums))
		for _, c := range tc.nums {
			name = append(name, strconv.Itoa(c))
		}

		t.Run(strings.Join(name, " "), func(t *testing.T) {
			result := Solution(tc.nums, tc.target)
			if result != tc.actual {
				t.Errorf("got %v, want %v", result, tc.actual)
			}
		})
	}
}
