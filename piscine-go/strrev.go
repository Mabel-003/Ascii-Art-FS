package piscinego

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)

	for a := 0; a < length/2; a++ {
		b := length - 1 - a
		runes[a], runes[b] = runes[b], runes[a]
	}
	return string(runes)
}
