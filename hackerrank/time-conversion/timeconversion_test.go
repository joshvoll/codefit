package hackerrank

import "testing"

func TestConversion(t *testing.T) {
	input := "07:05:45PM"
	res := timeConversion(input)
	if res != "19:05:45" {
		t.Fatalf("want 19:05:45, got %s ", res)
	}
}

func TestConversionRecap(t *testing.T) {
	input := "07:05:45PM"
	res := timeConversionRecap(input)
	if res != "19:05:45" {
		t.Fatalf("want 19:05:45, got %s ", res)
	}
}
