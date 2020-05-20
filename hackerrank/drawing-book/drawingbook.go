package hackerrank

import "fmt"

func pageCount(n, p int32) int32 {
	pages := n / 2
	right := p / 2
	left := pages - right
	fmt.Println("Pages : ", pages, " right : ", right, " left : ", left)
	if right < left {
		return right
	}
	return left
}
