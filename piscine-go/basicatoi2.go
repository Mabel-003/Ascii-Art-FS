package piscinego

func BasicAtoi2(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		digit := int(r - '0')
		result = result*10 + digit
	}
	return result
}
