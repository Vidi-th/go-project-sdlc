package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{1, 2, 3},
		{-1, -1, -2},
		{1.5, 2.5, 4},
	}

	for _, tt := range tests {
		t.Run("Addition", func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{5, 3, 2},
		{-1, 1, -2},
		{2.5, 1.5, 1},
	}

	for _, tt := range tests {
		t.Run("Subtraction", func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{3, 4, 12},
		{-2, 5, -10},
		{1.5, 2, 3},
	}

	for _, tt := range tests {
		t.Run("Multiplication", func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, expected float64
		errExpected    bool
	}{
		{10, 2, 5, false},
		{5, 0, 0, true}, // Pembagian dengan nol, diharapkan ada error
		{7, 3, 2.333, false},
	}

	for _, tt := range tests {
		t.Run("Division", func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.errExpected && err == nil {
				t.Errorf("Divide(%f, %f) should return error but got result %f", tt.a, tt.b, result)
			} else if !tt.errExpected && err != nil {
				t.Errorf("Divide(%f, %f) returned an error: %v", tt.a, tt.b, err)
			} else if result != tt.expected {
				t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
