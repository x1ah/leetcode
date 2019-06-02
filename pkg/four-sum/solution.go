package four_sum

import "sort"

// https://leetcode.com/problems/4sum/
func threeSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			if end == i {
				break
			}
			tmpValue := nums[i] + nums[start] + nums[end]
			if tmpValue < target {
				start++
			} else if tmpValue == target {
				if i < start {
					res = append(res, []int{nums[i], nums[start], nums[end]})
				} else {
					res = append(res, []int{nums[start], nums[i], nums[end]})
				}
				start++
				end--

				for start < end && nums[start-1] == nums[start] && nums[end] == nums[end+1] {
					start++
					end--
				}

			} else {
				end--
			}
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := threeSum(nums[i+1:], target-nums[i])
		for _, k := range tmp {
			k = append([]int{nums[i]}, k...)
			res = append(res, k)
		}
	}

	return res
}
