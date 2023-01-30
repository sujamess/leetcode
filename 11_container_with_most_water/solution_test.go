package container_with_most_water

import (
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		height []int
		actual int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, actual: 49},
		{height: []int{1, 1}, actual: 1},
		{height: []int{4, 3, 2, 1, 4}, actual: 16},
		{height: []int{1, 2, 4, 3}, actual: 4},
	}

	for _, tc := range testCases {
		name := make([]string, 0, len(tc.height))
		for _, h := range tc.height {
			name = append(name, strconv.Itoa(h))
		}

		t.Run(strings.Join(name, " "), func(t *testing.T) {
			result := Solution(tc.height)
			if result != tc.actual {
				t.Errorf("got %d, want %d", result, tc.actual)
			}
		})
	}
}
