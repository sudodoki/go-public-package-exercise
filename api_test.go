package gopublicpackageexercise

import "testing"

func TestAddInt(t *testing.T) {
	given1 := 1
	given2 := 3
	expected := 4

	if res := Add(given1, given2); res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestAddFloat(t *testing.T) {
	given1 := 1.1
	given2 := 3.9
	expected := 5.0

	// TODO: probably should compare diff < epsilon
	if res := Add(given1, given2); res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
