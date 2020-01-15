package bubbleSort

import "testing"

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	data := []int{4, 3, 2, 1}

	bubbleSort(data)

	t.Log(data)
}

func bubbleSort(data []int) {
	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}