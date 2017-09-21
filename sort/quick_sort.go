package sort

// QuickSort sorts the input data using quick sort algorithm.
func QuickSort(data []int) {
	quickSort(0, len(data) - 1, data)
}

func quickSort(low int, high int, data[]int) {
	if low >= high {
		return
	}
	p := partition(low, high, data)
	quickSort(low, p - 1, data)
	quickSort(p + 1, high, data)
}

func partition(low int, high int, data []int) int {
	pivot := data[low]
	for low < high {
		for low < high && pivot <= data[high] {
			high--
		}
		data[low] = data[high]
		for low < high && pivot >= data[low] {
			low++
		}
		data[high] = data[low]
	}
	data[low] = pivot
	return low
}