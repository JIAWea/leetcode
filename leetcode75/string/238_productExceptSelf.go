package string

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。


示例 1:
输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：
2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内


进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
*/

func productExceptSelf1(nums []int) []int {
	// 利用索引左侧所有数字的乘积和右侧所有数字的乘积（即前缀与后缀）相乘得到答案。
	// 该空间复杂度为 O(N),使用了两个数组

	var (
		numLen = len(nums)
		answer = make([]int, numLen)
		l      = make([]int, numLen)
		r      = make([]int, numLen)
	)

	l[0] = 1
	for i := 1; i < numLen; i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[numLen-1] = 1
	for i := numLen - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < numLen; i++ {
		answer[i] = l[i] * r[i]
	}

	return answer
}

func productExceptSelf2(nums []int) []int {
	var (
		numLen = len(nums)
		answer = make([]int, numLen)
	)

	answer[0] = 1
	for i := 1; i < numLen; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	R := 1
	for i := numLen - 1; i >= 0; i-- {
		answer[i] *= R
		R *= nums[i]
	}

	return answer
}
