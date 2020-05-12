package hackerrank

import "testing"

func TestSockMerchant(t *testing.T) {
	tc := int32(9)
	tt := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	res := SockMerchant(tc, tt)
	if res != 3 {
		t.Fatalf("want 3, got %v ", res)
	}
}
