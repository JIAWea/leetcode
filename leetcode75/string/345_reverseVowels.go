package string

import "strings"

/*
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。


示例 1：
输入：s = "hello"
输出："holle"

示例 2：
输入：s = "leetcode"
输出："leotcede"


提示：
1 <= s.length <= 3 * 105
s 由 可打印的 ASCII 字符组成
*/

func reverseVowels(s string) string {
	alphaMap := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
	}

	list := []byte(s)
	sLen := len(list)
	endIdx := sLen - 1

	for i := 0; i < sLen && i < endIdx; i++ {
		if _, ok := alphaMap[strings.ToLower(string(list[i]))]; ok {
			for j := endIdx; j >= 0 && j > i; j-- {
				if _, ok := alphaMap[strings.ToLower(string(list[j]))]; ok {
					list[i], list[j] = list[j], list[i]
					endIdx = j - 1
					break
				}
			}
		}
	}

	return string(list)
}
