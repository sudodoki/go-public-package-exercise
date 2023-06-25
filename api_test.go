package gopublicpackageexercise

import "testing"

func TestAdd(t *testing.T) {
	given1 := 1
	given2 := 3
	expected := 4

	if res := Add(given1, given2); res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
