package two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	idxMap := map[int][]int{}
	for idx, v := range nums {
		if _, ok := idxMap[v]; ok {
			idxMap[v] = append(idxMap[v], idx)
		} else {
			idxMap[v] = []int{idx}
		}
	}
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for start < end {
		tmp := nums[start] + nums[end]
		if tmp == target {
			if nums[start] == nums[end] {
				return idxMap[nums[end]]
			} else {
				return []int{idxMap[nums[start]][0], idxMap[nums[end]][0]}

			}
		} else if tmp < target {
			start++
		} else {
			end--
		}
	}

	return []int{}
}
