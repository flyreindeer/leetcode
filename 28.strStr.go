func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}

	sIndex := 0
	pIndex := 0
	space := 0

	keys := make([]int, 256)

	for i := 0; i < 256; i++ {
		keys[i] = -1
	}

	for i := 0; i < len(needle); i++ {
		keys[needle[i]] = i
	}

	for sIndex < len(haystack) {
		if haystack[sIndex] == needle[pIndex] {
			sIndex++
			pIndex++
			if pIndex == len(needle) {
				return sIndex - pIndex
			}
		} else {
			pIndex = 0
			if space+len(needle) < len(haystack) {
				space += len(needle) - keys[haystack[space+len(needle)]]
			} else {
				return -1
			}
			sIndex = space
		}
	}

	return -1
}