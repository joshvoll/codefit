package hackerrank

// CompareTriplets function below
func CompareTriplets(a []int32, b []int32) []int32 {
	alice := int32(0)
	bob := int32(0)
	var total []int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else {
			bob++
		}
	}
	total = append(total, alice, bob)
	return total
}
