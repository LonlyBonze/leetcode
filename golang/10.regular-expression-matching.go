/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}

	pp := p[0]
	matchAny := pp == '.'

	if len(p) == 1 {
		if len(s) != len(p) {
			return false
		}
		if matchAny {
			return true
		}
		return pp == s[0]
	}

	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		if !matchAny && pp != s[0] {
			return false
		}
		return isMatch(s[1:], p[1:])
	}

	for si := 0; si < len(s); si++ {
		if isMatch(s[si:], p[2:]) {
			return true
		}

		if !matchAny && pp != s[si] {
			return false
		}
	}
	return isMatch("", p[2:])
}

// @lc code=end

