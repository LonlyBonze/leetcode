/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, num := range nums {
		if indexPre, bHas := numMap[target-num]; bHas {
			return []int{indexPre, index}
		}
		numMap[num] = index
	}
	return []int{}
}

// @lc code=end
