package mergesort

// mergeSort is the main method of the merge
// take the slice of int and len
// find the middle by dividing the array / 2
// creating left and right properties
// add the number to the left and right variables
func mergeSort(items []int) []int {
	var num = len(items)
	if num == 1 {
		return items
	}
	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

// merge rigtn and left
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	count := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[count] = left[0]
			left = left[1:]
		} else {
			result[count] = right[0]
			right = right[1:]
		}
		count++
	}
	for j := 0; j < len(left); j++ {
		result[count] = left[j]
		count++
	}
	for j := 0; j < len(right); j++ {
		result[count] = right[j]
		count++
	}
	return
}
