package sort

// BubbleSort sort the int slices in ascending order
func BubbleSort(data []int) {
	var flag bool
	for i := len(data); i > 0; i-- {
		flag = true
		for j := 0; j + 1 < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

