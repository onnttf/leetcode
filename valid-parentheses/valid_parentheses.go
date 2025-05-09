package valid_parentheses

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	right2Left := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	runeList := make([]rune, 0, len(s))

	for _, v := range s {
		left, ok := right2Left[v]

		if !ok {
			runeList = append(runeList, v)
			continue
		}

		if len(runeList) == 0 {
			return false
		}

		lastIndex := len(runeList) - 1
		if runeList[lastIndex] != left {
			return false
		}

		runeList = runeList[:lastIndex]
	}

	return len(runeList) == 0
}
