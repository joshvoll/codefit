package kata

import "testing"

func TestEvaporator(t *testing.T) {
	decorant := Evaporator(10, 10, 10)
	if decorant != 22 {
		t.Errorf("want 22, got %d ", decorant)
	}
}
