package competitive

import "fmt"

// sorting algorythm using bubble sort
// check if the current number is less tnat the next number j > j+1 if so, swap
func sorting(arr []int32) []int32 {
	for i := int32(0); i < int32(len(arr)); i++ {
		for j := int32(0); j < int32(len(arr))-1; j++ {
			fmt.Println("Arreglos: ", arr[j], " arreglos siguiente : ", arr[j+1])
			if arr[j] > arr[j+1] {
				fmt.Println("make sort")
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
