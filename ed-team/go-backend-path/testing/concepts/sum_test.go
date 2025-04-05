package main

import "testing"

func TestSum(t *testing.T) {
	wanted := 5
	got := Sum(1, 4)
	if got != wanted {
		t.Errorf("Sum(1, 2) = %d; want %v", got, wanted)
		t.Fail()
	}
}

func TestMultipleSum(t *testing.T) {
	wanted := 10
	got := MultipleSum(1, 2, 3, 4)
	if got != wanted {
		t.Errorf("MultipleSum(1, 2, 3, 4) = %d; want %v", got, wanted)
		t.Fail()
	}
}
