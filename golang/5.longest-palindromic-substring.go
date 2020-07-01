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
	var startPos, endPos, length int
	indexMap := map[rune][]int{}
	for index, r := range s {
		indexs, bHas := indexMap[r]
		if !bHas {
			indexs = []int{}
		}
		indexs = append(indexs, index)
		indexMap[r] = indexs
	}
	for index, r := range s {
		indexs, _ := indexMap[r]
		if len(indexs) < 2 {
			continue
		}
		newIndexs := []int{}
		for _, pos := range indexs {
			if pos <= index {
				continue
			}
			newIndexs = append(newIndexs, pos)
			isPalindrome := true
			for sp, ep := index+1, pos-1; sp < ep; sp, ep = sp+1, ep-1 {
				if s[sp] != s[ep] {
					isPalindrome = false
					break
				}
			}
			if isPalindrome {
				if length < pos-index+1 {
					length = pos - index + 1
					startPos = index
					endPos = pos
				}
			}
		}
	}
	return s[startPos : endPos+1]
}
// @lc code=end

