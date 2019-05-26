package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 只过非负，三个数一定需要包含复数
		if nums[i] > 0 {
			break
		} else if nums[i] == 0 {
			// 三个 0 的情况
			if len(nums) >= i+3 && nums[i+1] == 0 && nums[i+2] == 0 {
				res = append(res, []int{0, 0, 0})
			}
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			if end == i {
				break
			}
			tmpValue := nums[i] + nums[start] + nums[end]
			if tmpValue < 0 {
				start++
			} else if tmpValue == 0 {
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
