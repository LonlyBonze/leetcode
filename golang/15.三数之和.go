/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	ret := [][]int{}

	numMap := map[int][]int{}
	for index, num := range nums {
		numMap[num] = append(numMap[num], index)
		if num == 0 && len(numMap[num]) == 3 {
			ret = append(ret, []int{num, num, num})
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if i != numMap[nums[i]][0] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j != numMap[nums[j]][0] {
				continue
			}

			sum := nums[i] + nums[j]
			pos, ok := numMap[-sum]
			if !ok {
				continue
			}

			if (-sum == nums[i] || -sum == nums[j]) && len(pos) <= 1 {
				continue
			}

			pNum := 0
			preNum := 0

			for _, p := range pos {
				if p > j {
					pNum++
				} else {
					preNum++
				}
			}

			if -sum != nums[i] && -sum != nums[j] && (pNum == 0 || preNum > 0) {
				continue
			}

			ret = append(ret, []int{nums[i], nums[j], -sum})
		}
	}

	return ret
}

// @lc code=end

