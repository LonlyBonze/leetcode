/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	sList := strings.Split(s, "")
	s = " " + strings.Join(sList, " ") + " "
	var startPos, endPos, length int
	for i := 0; i < len(s); i++ {
		var si, ei int
		for si, ei = i-1, i+1; si >= 0 && ei < len(s); si, ei = si-1, ei+1 {
			if s[si] != s[ei] {
				break
			}
		}
		si++
		ei--
		if length < ei-si+1 {
			length = ei - si + 1
			startPos = si
			endPos = ei
		}
	}
	s = s[startPos : endPos+1]
	sList = strings.Split(s, " ")
	s = strings.Join(sList, "")
	return s
}
// @lc code=end

