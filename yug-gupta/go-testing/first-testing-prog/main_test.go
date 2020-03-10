package main

import "testing"

func TestIncrement(t *testing.T) {
	if increment(2) != 4 {
		t.Error("Wrong answer;Expected 4")
	}
}
