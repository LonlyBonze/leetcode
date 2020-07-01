/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	pos := map[rune]int{}
	start := 0
	for i, c := range s {
		p, bHas := pos[c]
		if bHas && p >= start {
			length := i - start
			start = p + 1
			if length > maxLength {
				maxLength = length
			}
		}
		pos[c] = i
	}
	length := len(s) - start
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
// @lc code=end

