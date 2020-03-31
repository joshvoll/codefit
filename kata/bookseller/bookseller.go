package kata

import (
	"fmt"
	"strconv"
	"strings"
)

// StockList get to list return a string
func StockList(listArt []string, listCart []string) string {
	if len(listArt) == 0 || len(listCart) == 0 {
		return ""
	}
	var cats []string
	for _, c := range listCart {
		sum := 0
		for _, s := range listArt {
			if s[0:1] == c {
				if s, err := strconv.Atoi(strings.Split(s, " ")[1]); err == nil {
					sum += s
				}
			}
		}
		cats = append(cats, fmt.Sprintf(`(%s : %d)`, c, sum))
	}
	return strings.Join(cats, " - ")

}
