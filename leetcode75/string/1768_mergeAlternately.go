package string

func mergeAlternately1(word1, word2 string) string {
	w1Len := len(word1)
	w2Len := len(word2)

	maxLen := w1Len
	if w2Len > w1Len {
		maxLen = w2Len
	}

	var ret string

	for i := 0; i < maxLen; i++ {
		if i < w1Len {
			ret += string(word1[i])
		}

		if i < w2Len {
			ret += string(word2[i])
		}
	}

	return ret
}

func mergeAlternately2(word1, word2 string) string {
	w1Len := len(word1)
	w2Len := len(word2)

	var ret string

	for i := 0; i < w1Len; i++ {
		ret += string(word1[i])

		if i < w2Len {
			ret += string(word2[i])
		}
	}

	if w1Len < w2Len {
		ret += word2[w1Len:]
	}

	return ret
}

func mergeAlternately3(word1, word2 string) string {
	w1Len := len(word1)
	w2Len := len(word2)

	ret := make([]byte, 0, w1Len+w2Len)

	for i := 0; i < w1Len; i++ {
		ret = append(ret, word1[i])

		if i < w2Len {
			ret = append(ret, word2[i])
		}
	}

	if w1Len < w2Len {
		ret = append(ret, word2[w1Len:]...)
	}

	return string(ret)
}
