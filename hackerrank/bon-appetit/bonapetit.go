package hackerrank

import "fmt"

func bonAppetit(bill []int32, k, b int32) {
	var sum int32
	var res int32

	for i := int32(0); i < int32(len(bill)); i++ {
		if i != k {
			sum += bill[i]
		}
	}
	res = b - (sum / 2)
	if res == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(res)
	}
}
