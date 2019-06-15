package next_permutation

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	cursor := len(nums) - 2
	for cursor >= 0 && nums[cursor] >= nums[cursor+1] {
		cursor--
	}

	if cursor < 0 {
		reverse(nums, 0)
		return
	}

	cursor2 := len(nums) - 1
	for cursor2 >= 0 && nums[cursor2] <= nums[cursor] {
		cursor2--
	}

	if cursor2 < cursor {
		return
	}

	nums[cursor], nums[cursor2] = nums[cursor2], nums[cursor]
	reverse(nums, cursor+1)
}

func reverse(nums []int, start int) {
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
