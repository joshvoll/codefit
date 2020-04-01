package hackerrank

import "fmt"

// PlusMinus method implementation
func PlusMinus(arr []int32) {
	positive, negative, zero := int32(0), int32(0), int32(0)
	for _, v := range arr {
		if v < 0 {
			negative++
		} else if v > 0 {
			positive++
		} else {
			zero++
		}
	}
	lens := len(arr)
	fmt.Println("Positive: ", float64(positive)/float64(lens))
	fmt.Println("Negative: ", float64(negative)/float64(lens))
	fmt.Println("zero: ", float64(zero)/float64(lens))
}

// plusMinus just a for loop test
func plusMinus(arr []int32) {
	positive, negative, zero := int32(0), int32(0), int32(0)
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative++
		} else if arr[i] > 0 {
			positive++
		} else {
			zero++
		}
	}
	fmt.Println(positive, negative, zero)
}
