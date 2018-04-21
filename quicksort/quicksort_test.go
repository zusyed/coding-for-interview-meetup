package quicksort_test

import (
	"testing"

	"github.com/zusyed/interview-meetup/quicksort"
	"github.com/zusyed/interview-meetup/utils"
)

func TestSort(t *testing.T) {
	tests := map[string]struct {
		a        []int
		expected []int
	}{
		"1": {
			a:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		"2": {
			a:        []int{3, 2, 4, 6, 1, 5, 8, 9, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for k, v := range tests {
		t.Logf("Running test: %s", k)
		quicksort.Sort(v.a)
		if !utils.ArraysEqual(v.a, v.expected) {
			t.Errorf("Expected result to be %t but was %t", v.expected, v.a)
		}
	}
}
