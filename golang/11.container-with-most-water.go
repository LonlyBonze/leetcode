/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	f := 0
	b := len(height) - 1

	area := 0
	for f < b {
		var a int
		if height[f] < height[b] {
			a = (b - f) * height[f]
			f++
		} else {
			a = (b - f) * height[b]
			b--
		}
		if a > area {
			area = a
		}
	}
	return area
}

// @lc code=end

