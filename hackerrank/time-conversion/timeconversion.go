package hackerrank

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	//AM, PM := "AM", "PM"
	const inputLayout = "03:04:05PM"
	const outputLayout = "15:04:05"
	t, err := time.Parse(inputLayout, s)
	if err != nil {
		fmt.Println(err)
	}
	res := t.Format(outputLayout)
	return res
}
