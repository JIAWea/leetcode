package string

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，
能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。


示例 1：
输入：flowerbed = [1,0,0,0,1], n = 1
输出：true

示例 2：
输入：flowerbed = [1,0,0,0,1], n = 2
输出：false


提示：
1 <= flowerbed.length <= 2 * 104
flowerbed[i] 为 0 或 1
flowerbed 中不存在相邻的两朵花
0 <= n <= flowerbed.length
*/

func canPlaceFlowers1(flowerbed []int, n int) bool {
	num := 0
	l := len(flowerbed)
	endIdx := l - 1

	for i := 0; i < l; i++ {
		preIdx := i - 1
		nextIdx := i + 1

		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 {
			if flowerbed[i] == 0 && nextIdx < l && flowerbed[nextIdx] == 0 {
				flowerbed[i] = 1
				num += 1
				continue
			}
			if i == endIdx {
				flowerbed[i] = 1
				num += 1
				continue
			}
			continue
		}

		if flowerbed[preIdx] == 0 && nextIdx < l && flowerbed[nextIdx] == 0 {
			flowerbed[i] = 1
			num += 1
			continue
		}

		if i == endIdx && flowerbed[preIdx] == 0 {
			flowerbed[i] = 1
			num += 1
			continue
		}
	}

	if num >= n {
		return true
	}

	return false
}

/**
* 贪心
* 能种花的地方
*  当前位置没有花
*  前面要么没有花, 要么是边界
*  后面要么没有花, 要么是边界
 */
func canPlaceFlowers2(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			n--
			flowerbed[i] = 1
		}
		if n <= 0 {
			return true
		}
	}
	return false
}
