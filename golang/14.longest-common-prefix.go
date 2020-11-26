/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	p := 0
	for ; p < len(strs[0]); p++ {
		for i := 1; i < len(strs); i++ {
			if p >= len(strs[i]) || strs[i][p] != strs[0][p] {
				return strs[0][:p]
			}
		}
	}
	return strs[0][:p]
}

// @lc code=end

