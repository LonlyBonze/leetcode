/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		r := x % 10
		x = x / 10
		if x == y {
			return true
		}
		y = y*10 + r
	}
	if x == y {
		return true
	}

	return false
}

// @lc code=end

