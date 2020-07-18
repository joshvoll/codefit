package mergesort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	slice := generateSlice(20)
	res := mergeSort(slice)
	fmt.Println("\n--- Unosrted--- \n\n", slice)
	fmt.Println(res)

}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
