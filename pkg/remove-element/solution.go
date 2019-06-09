package remove_element

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] != val {
			start++
			continue
		}

		for nums[end] == val {
			if start >= end {
				return start
			}
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
	}

	return start
}
