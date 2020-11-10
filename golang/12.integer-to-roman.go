/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	var roman string

	if num >= 1000 {
		for i := 0; i < num/1000; i++ {
			roman += "M"
		}
		num = num % 1000
	}
	if num >= 900 {
		roman += "CM"
		num -= 900
	}
	if num >= 500 {
		roman += "D"
		num -= 500
	}
	if num >= 400 {
		roman += "CD"
		num -= 400
	}
	if num >= 100 {
		for i := 0; i < num/100; i++ {
			roman += "C"
		}
		num = num % 100
	}
	if num >= 90 {
		roman += "XC"
		num -= 90
	}
	if num >= 50 {
		roman += "L"
		num -= 50
	}
	if num >= 40 {
		roman += "XL"
		num -= 40
	}
	if num >= 10 {
		for i := 0; i < num/10; i++ {
			roman += "X"
		}
		num = num % 10
	}
	if num >= 9 {
		roman += "IX"
		num -= 9
	}
	if num >= 5 {
		roman += "V"
		num -= 5
	}
	if num >= 4 {
		roman += "IV"
		num -= 4
	}
	for i := 0; i < num; i++ {
		roman += "I"
	}

	return roman
}

// @lc code=end

