package hackerrank

import "fmt"

// birthdayCakeCandlesRecap it just another test for study
func birthdayCakeCandlesRecap(arr []int32) int32 {
	max := int32(0)
	count := 0
	for i := int32(0); i < int32(len(arr)); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	for j := int32(0); j < int32(len(arr)); j++ {
		if arr[j] == max {
			count++
		}
	}
	return int32(count)
}

func birthdayCakeCandlesRecap2(arr []int32) int32 {
	max := int32(0)
	count := 0
	for i := int32(0); i < int32(len(arr)); i++ {
		fmt.Println("ARR : ", arr[i], " MAX ", max, " COUNT ", count)
		if arr[i] == max {
			count++
		}
		if max < arr[i] {
			max = arr[i]
			count = 1
		}

	}
	return int32(count)
}
