/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	num := 0
	pre := 0
	for _, r := range s {
		var i int
		switch r {
		case 'I':
			i = 1
		case 'V':
			i = 5
		case 'X':
			i = 10
		case 'L':
			i = 50
		case 'C':
			i = 100
		case 'D':
			i = 500
		case 'M':
			i = 1000
		}
		num += i
		if pre < i {
			num -= 2 * pre
		}
		pre = i
	}
	return num
}

// @lc code=end

