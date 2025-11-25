package piscinego

func Swap(a *int, b *int) {
	i := *a
	*a = *b
	*b = i
}
