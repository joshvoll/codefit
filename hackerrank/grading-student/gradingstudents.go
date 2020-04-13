package hackerrank

func gradingStudents(grades []int32) []int32 {
	res := grades
	for i, v := range grades {
		if v >= 38 && res[i]%5 > 2 {
			res[i] += 5 - res[i]%5
		}
	}
	return res
}
