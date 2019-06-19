package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	for idx, v := range nums {
		if v == target {
			return idx
		}
	}

	return -1
}
