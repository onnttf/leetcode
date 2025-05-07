package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	firstString := strs[0]
	for i := 0; i < len(firstString); i++ {
		prefix := firstString[i]
		for j := 1; j < len(strs); j++ {
			nextString := strs[j]
			if i >= len(nextString) {
				return firstString[:i]
			}
			nextPrefix := nextString[i]
			if prefix != nextPrefix {
				return firstString[:i]
			}
		}
	}
	return firstString
}
