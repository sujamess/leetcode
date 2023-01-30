package combination_sum

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		candidates []int
		target     int
		actual     [][]int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7, actual: [][]int{{2, 2, 3}, {7}}},
		{candidates: []int{2, 3, 5}, target: 8, actual: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{candidates: []int{2}, target: 1, actual: [][]int{}},
	}

	for _, tc := range testCases {
		name := make([]string, 0, len(tc.candidates))
		for _, c := range tc.candidates {
			name = append(name, strconv.Itoa(c))
		}

		t.Run(strings.Join(name, " "), func(t *testing.T) {
			result := Solution(tc.candidates, tc.target)
			if !reflect.DeepEqual(result, tc.actual) {
				t.Errorf("got %v, want %v", result, tc.actual)
			}
		})
	}
}
