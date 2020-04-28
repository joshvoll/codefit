package hackerrank

func gradingStudentsRecap(ar []int32) []int32 {
	res := ar
	for i := int32(0); i < int32(len(ar)); i++ {
		if ar[i] >= 38 && ar[i]%5 > 2 {
			res[i] += 5 - res[i]%5
		}
	}
	return res
}
