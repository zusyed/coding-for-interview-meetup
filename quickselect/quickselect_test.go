package quickselect_test

import (
	"testing"

	"github.com/zusyed/interview-meetup/quickselect"
)

func TestSort(t *testing.T) {
	tests := map[string]struct {
		a        []int
		k        int
		expected int
	}{
		"1": {
			a:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:        0,
			expected: 1,
		},
		"2": {
			a:        []int{3, 2, 4, 6, 1, 5, 8, 9, 7},
			k:        8,
			expected: 9,
		},
		"3": {
			a:        []int{3, 2, 4, 6, 1, 5, 8, 9, 7},
			k:        3,
			expected: 6,
		},
	}

	for k, v := range tests {
		t.Logf("Running test: %s", k)
		actual := quickselect.Select(v.a, v.k)
		if actual != v.expected {
			t.Errorf("Expected result to be %d but was %d", v.expected, actual)
		}
	}
}
