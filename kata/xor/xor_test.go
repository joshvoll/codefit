package kata

import "testing"

func Test_xor(t *testing.T) {
	res := Xor(false, false)
	if res != false {
		t.Fatalf("shoudl get false : %v ", res)
	}
}
