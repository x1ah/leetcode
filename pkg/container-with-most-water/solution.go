package container_with_most_water

// https://leetcode.com/problems/container-with-most-water
// 从最矮的边，每次加一个单位的高度，画一条水平线，算与该水平线相交的最大面积（第一次相交于最后一次相交）
// 直到最高边
func maxArea(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}

	min, max := MinAndMax(height)

	if min == max {
		return min * (length - 1)
	}

	maxarea := 0
	for i := min; i <= max; i++ {
		cursor := 0
		for j := cursor; j < length; j++ {
			if height[j] >= i {
				cursor = j
				break
			}
		}
		if cursor == length-1 {
			return maxarea
		}

		for j := length - 1; j > cursor; j-- {
			if height[j] >= i {
				area := (j - cursor) * i
				if maxarea < area {
					maxarea = area
				}
				break
			} else if j == cursor+1 {
				return maxarea
			}
		}
	}

	return maxarea
}

func MinAndMax(values []int) (min, max int) {
	min, max = values[0], values[0]
	for _, v := range values {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	return
}
