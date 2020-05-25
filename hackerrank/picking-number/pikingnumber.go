package hackerrank

func pickingNumer(a []int32) int32 {
	var mv = make(map[int32]int32)
	for i := 0; i < len(a); i++ {
		mv[a[i]]++
	}
	var max int32 = 0
	for k, v := range mv {
		if max < mv[k+1]+v {
			max = mv[k+1] + v
		}

		if max < mv[k-1]+v {
			max = mv[k-1] + v
		}
	}
	return max
}
