package q1

import "testing"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	d := map[int] int {}
	for i := 0; i < length; i++ {
		part := target - nums[i]
		if v, ok := d[nums[i]]; ok {
			return []int{v, i}
		} else {
			d[part] = i
		}
	}
	return nil
}

func TestQ1(t *testing.T)  {
	data := []int{2, 7, 11, 15}
	t.Log(twoSum(data, 9))
}
