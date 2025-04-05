package main

import "testing"

type table struct {
	name         string
	a, b, wanted int
}

func TestMultiply(t *testing.T) {
	table := []table{
		{"3x1", 3, 1, 3},
		{"3x2", 3, 2, 6},
		{"3x3", 3, 3, 9},
		{"3x4", 3, 4, 12},
		{"3x5", 3, 5, 15},
		{"3x6", 3, 6, 18},
		{"3x7", 3, 7, 21},
		{"3x8", 3, 8, 24},
		{"3x9", 3, 9, 27},
		{"3x10", 3, 10, 30},
	}
	for _, row := range table {
		t.Run(row.name, func(t *testing.T) {
			got := multiply(row.a, row.b)
			if got != row.wanted {
				t.Errorf("multiply(%v, %v) = %d, want %d)", row.a, row.b, got, row.wanted)
				t.Fail()
			}
		})
	}
}
