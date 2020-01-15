package selectOrder

import "testing"

// 选择排序
func findMax(data []int) (index int, value int) {
	maxIndex := 0
	maxValue := data[maxIndex]

	for i := 1; i < len(data); i++ {
		if data[i] > maxValue {
			maxIndex = i
			maxValue = data[i]
		}
	}

	return maxIndex, maxValue
}

func TestSelectOrder(t *testing.T) {
	data := []int{2, 3, 1, 5}
	result := make([]int, 4)
	dataLen := len(data)

	for i := 0; i < dataLen; i++ {
		maxIndex, maxValue := findMax(data)
		result[i] = maxValue
		data = append(data[:maxIndex], data[maxIndex+1:]...)
	}

	t.Log(result)
}

