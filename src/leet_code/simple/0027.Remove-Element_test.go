package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {

	arr := []int{1, 1, 2, 2, 3, 3}
	len := removeElement(arr, 1)
	println(len)

}
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
