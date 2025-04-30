package roman_to_integer

func romanToInt(s string) int {
	roman2Value := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentValue := roman2Value[s[i]]

		if i+1 < n && currentValue < roman2Value[s[i+1]] {
			total -= currentValue
		} else {
			total += currentValue
		}
	}

	return total
}
