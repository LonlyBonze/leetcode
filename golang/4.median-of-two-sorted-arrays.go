/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i1s, i1e, i2s, i2e int
	i1e = len(nums1) - 1
	i2e = len(nums2) - 1
	var a, b int
	b = 1
	for a < b {
		if i1s > i1e && i2s > i2e {
			break
		} else if i1s > i1e {
			a = nums2[i2s]
			b = nums2[i2e]
			i2s++
			i2e--
		} else if i2s > i2e {
			a = nums1[i1s]
			b = nums1[i1e]
			i1s++
			i1e--
		} else {
			if nums1[i1s] > nums2[i2s] {
				a = nums2[i2s]
				i2s++
			} else {
				a = nums1[i1s]
				i1s++
			}
			if nums1[i1e] > nums2[i2e] {
				b = nums1[i1e]
				i1e--
			} else {
				b = nums2[i2e]
				i2e--
			}
		}
	}
	return float64(a+b) / 2
}
// @lc code=end

