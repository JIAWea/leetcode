package string

import (
	"bytes"
)

/*
1071. 字符串的最大公因子
提示：对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。


示例 1：
输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"

示例 2：
输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"

示例 3：
输入：str1 = "LEET", str2 = "CODE"
输出：""
*/

func checkGcd(x, str string) bool {
	var ans bytes.Buffer
	for i := 0; i < len(str)/len(x); i++ {
		ans.WriteString(x)
	}
	return ans.String() == str
}

func gcdOfStrings(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	min := len1
	if len2 < len1 {
		min = len2
	}

	for i := min; i > 0; i-- {
		if len1%i == 0 && len2%i == 0 {
			x := str1[:i]
			if checkGcd(x, str1) && checkGcd(x, str2) {
				return x
			}
		}
	}

	return ""
}
