package search_insert_position

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	length := len(nums)
	if nums[length-1] < target {
		return length
	}

	if nums[0] > target {
		return 0
	}

	l, r := 0, length-1
	middle := (l + r) / 2
	for l <= r {
		middle = (l + r) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}

	if nums[middle] < target {
		return middle + 1
	} else {
		return middle
	}
}
