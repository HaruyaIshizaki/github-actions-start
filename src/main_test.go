package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(1)
	if result != "odd" {
		t.Errorf("Expected odd, got %s", result)
	}
}

func TestEvenOrOdd2(t *testing.T) {
	result := EvenOrOdd(2)
	if result != "even" {
		t.Errorf("Expected even, got %s", result)
	}
}