package composition_test

import (
	"main/composition"
	"testing"
)

func TestSquare(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{6, 36},
		{0, 0},
		{-3, 9},
		{10, 100},
	}

	for _, tc := range testCases {
		got := composition.Square(tc.input)
		if got != tc.expected {
			t.Errorf("Square(%d) = %d; want %d", tc.input, got, tc.expected)
		}
	}
}

func TestInc(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{6, 7},
		{0, 1},
		{-3, -2},
	}

	for _, tc := range testCases {
		got := composition.Inc(tc.input)
		if got != tc.expected {
			t.Errorf("Inc(%d) = %d; want %d", tc.input, got, tc.expected)
		}
	}
}

func TestCompose(t *testing.T) {
	h := composition.Compose(composition.Square, composition.Inc) // h(x) = Square(Inc(x))

	testCases := []struct {
		input    int
		expected int
	}{
		{6, 49},   // (6+1)^2 = 49
		{0, 1},    // (0+1)^2 = 1
		{-3, 4},   // (-3+1)^2 = 4
		{10, 121}, // (10+1)^2 = 121
	}

	for _, tc := range testCases {
		got := h(tc.input)
		if got != tc.expected {
			t.Errorf("Compose(%d) = %d; want %d", tc.input, got, tc.expected)
		}
	}
}
