package hackerrank

import "fmt"

// CompareTripletsRecap is just a recap method
func CompareTripletsRecap(a, b []int32) []int32 {
	alice := int32(0)
	bob := int32(0)
	fmt.Println(a)
	fmt.Println(b)
	for i := int32(0); i < int32(len(a)); i++ {
		if a[i] > b[i] {
			alice++
			fmt.Println(a[i], " > ", b[i], " Alice get a point")
		}
		if a[i] < b[i] {
			fmt.Println(a[i], " < ", b[i], " bob get a point")
			bob++
		}
	}
	return []int32{alice, bob}
}
