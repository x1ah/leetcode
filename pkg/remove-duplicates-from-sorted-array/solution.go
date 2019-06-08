package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[length-1] {
			continue
		}

		nums[length], nums[i] = nums[i], nums[length]
		length++
	}

	return length
}
