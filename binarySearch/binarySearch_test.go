package binarySearch

import (
	"testing"
)

var (
	count = 100000000
)

/**
 * 二分查找
 */
func BinarySearch(data []int, target int) (index int) {
	low := 0
	height := len(data) - 1

	for ; low < height;  {
		mid := int((low + height) / 2)
		guess := data[mid]

		if guess == target  {
			return mid
		}

		if guess > target {
			height = mid - 1
		} else {
			low = mid + 1
		}
	}

	return
}

func TestBinarySearch(t *testing.T) {
	var data = make([]int, count)

	for i := 0; i < count; i++ {
		data[i] = i+1
	}

	target := 2

	index := BinarySearch(data, target)

	t.Log(index)
}