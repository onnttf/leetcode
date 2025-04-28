package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	y := 0

	for x > y {
		y = y*10 + x%10
		x = x / 10
	}

	return x == y || x == y/10
}
