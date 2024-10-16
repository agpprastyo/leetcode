func myAtoi(s string) int {
	const (
		INT_MAX = 1<<31 - 1 // 2^31 - 1
		INT_MIN = -1 << 31   // -2^31
	)

	i := 0
	n := len(s)

	// Step 1: Ignore leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	// Step 2: Determine the sign
	sign := 1
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3: Convert the digits
	result := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Check for overflow and underflow
		if result > (INT_MAX-digit)/10 {
			if sign == 1 {
				return INT_MAX
			} else {
				return INT_MIN
			}
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}
