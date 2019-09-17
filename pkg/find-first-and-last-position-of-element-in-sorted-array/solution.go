package find_first_and_last_position_of_element_in_sorted_array

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := -1, -1
	for idx, v := range nums {
		if v == target {
			if start == -1 {
				start = idx
				end = idx
			}

			if end != -1 {
				end = idx
			}
		}
	}
	return []int{start, end}
}
