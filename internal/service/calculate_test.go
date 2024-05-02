package service

import "testing"

func TestCalculateFactorials(t *testing.T) {
	c := &Calculator{}

	//Good
	a, b := 3, 5
	factA, factB := c.CalculateFactorials(a, b)
	if factA != 6 || factB != 120 {
		t.Errorf("Expected factorials: %d and %d, got: %d and %d", 6, 120, factA, factB)
	}

	a, b = 1, 2
	factA, factB = c.CalculateFactorials(a, b)
	if factA != 1 || factB != 2 {
		t.Errorf("Expected factorials: %d and %d, got: %d and %d", 1, 2, factA, factB)
	}

	// Bad
	// 0
	a, b = 0, 5
	factA, factB = c.CalculateFactorials(a, b)
	if factA != 1 || factB != 120 {
		t.Errorf("Expected factorials: %d and %d, got: %d and %d", 1, 120, factA, factB)
	}

	a, b = 3, 0
	factA, factB = c.CalculateFactorials(a, b)
	if factA != 6 || factB != 1 {
		t.Errorf("Expected factorials: %d and %d, got: %d and %d", 120, 1, factA, factB)
	}

	//negative numbers
	a, b = -3, -5
	factA, factB = c.CalculateFactorials(a, b)
	if factA != 1 || factB != 1 {
		t.Errorf("Expected factorials: %d and %d, got: %d and %d", 1, 1, factA, factB)
	}
}
