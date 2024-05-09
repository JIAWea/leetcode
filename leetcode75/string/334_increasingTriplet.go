package string

import "math"

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。


示例 1：
输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意

示例 2：
输入：nums = [5,4,3,2,1]
输出：false
解释：不存在满足题意的三元组

示例 3：
输入：nums = [2,1,5,0,4,6]
输出：true
解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6


提示：
1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1


进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
*/

/*
双向遍历，维护两个数组，只要满足 leftMin[i-1] < nums[i] < rightMax[i+1] 即可
*/
func increasingTriplet1(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
贪心

	赋初始值的时候，已经满足second > first了，现在找第三个数third

(1) 如果third比second大，那就是找到了，直接返回true
(2) 如果third比second小，但是比first大，那就把second指向third，然后继续遍历找third
(3) 如果third比first还小，那就把first指向third，然后继续遍历找third（这样的话first会跑到second的后边，但是不要紧，因为在second的前边，老first还是满足的）
*/
func increasingTriplet2(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
