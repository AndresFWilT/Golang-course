package main

import "testing"

const (
	fiboSequence int = 12
	wanted       int = 144
)

func TestFibonacciFor(t *testing.T) {
	got := fibonacciFor(fiboSequence)
	t.Logf("fibonacciFor(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciFor(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func TestFibonacciRec(t *testing.T) {
	got := fibonacciRec(fiboSequence)
	t.Logf("fibonacciRec(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciRec(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func TestFibonacciRecMem(t *testing.T) {
	got := fibonacciRecMem(fiboSequence)
	t.Logf("fibonacciRecMem(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciRecMec(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciFor(30)
	}
}

func BenchmarkFibonacciRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRec(30)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecMem(30)
	}
}
