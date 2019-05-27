package three_sum_closest

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/
// 遍历法
func threeSumClosest(nums []int, target int) int {
	res := target
	sort.Ints(nums)
	miss, isInit := 0, true

	for i := 0; i < len(nums); i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			if end == i {
				break
			}
			tmpValue := nums[i] + nums[start] + nums[end]
			if tmpValue == target {
				return target
			}
			v := int(math.Abs(float64(tmpValue - target)))
			if isInit {
				miss, isInit, res = v, false, tmpValue
			} else if v < miss {
				miss, res = v, tmpValue
			}
			if tmpValue < target {
				start++
			} else {
				end--
			}
		}
	}

	return res
}
