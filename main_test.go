package main

import "testing"

func TestAdd(t *testing.T) {
	s := add(10, 20)
	if s != 30 {
		t.Errorf("Expected: 30, Actual: %d", s)
	}
}
