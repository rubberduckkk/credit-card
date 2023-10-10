package luhn

func IsValidLuhn(digits []int) bool {
	var sum int
	for i := len(digits) - 2; i >= 0; i-- {
		if i%2 == 0 {
			sum += digits[i] * 2
			if digits[i] > 4 {
				sum -= 9
			}
		} else {
			sum += digits[i]
		}
	}
	return 10-sum%10 == digits[len(digits)-1]
}
