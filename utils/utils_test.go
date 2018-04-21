package utils_test

import (
	"testing"

	"github.com/zusyed/interview-meetup/utils"
)

func TestArraysEqual(t *testing.T) {
	tests := map[string]struct {
		a1       []int
		a2       []int
		expected bool
	}{
		"1": {
			a1:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			a2:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: true,
		},
		"2": {
			a1:       []int{1, 2, 3, 4, 5, 6, 7, 8},
			a2:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: false,
		},
		"3": {
			a1:       []int{1, 2, 3, 4, 5, 6, 7, 8, 1},
			a2:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: false,
		},
	}

	for k, v := range tests {
		t.Logf("Running test: %s", k)
		actual := utils.ArraysEqual(v.a1, v.a2)
		if v.expected != actual {
			t.Errorf("Expected result to be %t but was %t", v.expected, actual)
		}
	}
}
