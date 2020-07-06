/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	} else if numRows == 1 {
		return s
	}
	step := numRows<<1 - 2
	hStep := step / 2
	subList := make([]strings.Builder, hStep+1)
	for i, r := range s {
		si := i % step
		var sr int
		if si < hStep {
			sr = si % hStep
		} else {
			sr = hStep - si%hStep
		}
		if sr == 0 {
			subList[si].WriteRune(r)
		} else {
			subList[sr].WriteRune(r)
		}
	}
	result := strings.Builder{}
	for _, sub := range subList {
		result.WriteString(sub.String())
	}
	return result.String()
}
// @lc code=end

