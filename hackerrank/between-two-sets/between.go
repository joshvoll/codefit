package hackerrrank

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	c := int32(0)
	for i := int32(1); i <= 100; i++ {
		factor := true
		for j := int32(0); j < int32(len(a)); j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}
		if factor {
			for j := int32(0); j < int32(len(b)); j++ {
				if b[j]%i != 0 {
					factor = false
					break
				}
			}
		}
		if factor {
			c++
		}
	}
	fmt.Println(c)
	return c
}
