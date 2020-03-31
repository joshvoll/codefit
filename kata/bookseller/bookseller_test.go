package kata

import "testing"

func TestBookSeller(t *testing.T) {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	str := StockList(b, c)
	if str != "(A : 0) - (B : 1290) - (C : 515) - (D : 600)" {
		t.Fatalf("want `(A : 0) - (B : 1290) - (C : 515) - (D : 600)`, got : %s ", str)
	}

}
