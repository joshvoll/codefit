package hackerrank

import "fmt"

func plusMinusRecap(ar []int32) {
	var positive, negative, zero int32
	for i := int32(0); i < int32(len(ar)); i++ {
		if int32(ar[i]) < 0 {
			negative++
		} else if int32(ar[i]) > 0 {
			positive++
		}
		zero++
	}
	lens := int32(len(ar))
	fmt.Println("recap negative: ", float64(negative)/float64(lens))
}
