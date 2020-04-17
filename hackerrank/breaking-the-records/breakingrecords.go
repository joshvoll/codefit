package hackerrank

func breakingRecords(scores []int32) []int32 {
	var min, max, cmin, cmax int32
	for i := int32(0); i < int32(len(scores)); i++ {
		if i == 0 {
			min = scores[i]
			max = scores[i]
		} else {
			if scores[i] < min {
				cmin++
				min = scores[i]
			}
			if scores[i] > max {
				cmax++
				max = scores[i]
			}
		}
	}
	return []int32{cmax, cmin}
}
