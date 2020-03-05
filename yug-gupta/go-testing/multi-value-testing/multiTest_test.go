package main

import (
	"fmt"
	"testing"
)

func TestIncrement(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{2, 4},
		{4, 6},
		{-12, -10},
		{9090999, 896689},
	}
	for _, test := range tests {
		if Increment(test.input) != test.expected {
			t.Error("Test case failed. Output of:", test.input, " is expected to be", Increment(test.input), "and we received", test.expected)
		} else {
			fmt.Println(test.input, test.expected, "Passed")
		}
	}

}
