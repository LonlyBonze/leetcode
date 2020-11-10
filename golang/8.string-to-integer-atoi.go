/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	atoiMap := map[rune]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'0': 0,
	}

	strRune := []rune(str)

	var result int
	isNegative := false

	index := 0
	for ; index < len(strRune); index++ {
		if unicode.IsSpace(strRune[index]) {
			continue
		}

		if strRune[index] == '-' {
			isNegative = true
			index++
			break
		}
		if strRune[index] == '+' {
			index++
			break
		}
		_, bHas := atoiMap[strRune[index]]
		if bHas {
			break
		}
		return result
	}

	for ; index < len(strRune); index++ {
		value, bHas := atoiMap[strRune[index]]
		if !bHas {
			break
		}
		result = result*10 + value

		if result >= 2147483648 {
			break
		}
	}

	if isNegative {
		result = -result
	}
	if result > 2147483647 {
		return 2147483647
	} else if result < -2147483648 {
		return -2147483648
	}

	return result
}

// @lc code=end

