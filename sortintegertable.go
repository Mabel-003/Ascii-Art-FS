package piscinego

func SortIntegerTable(table []int) {
	n := len(table)
	for a := 0; a < n-1; a++ {
		for b := 0; b < n-1-a; b++ {
			if table[b] > table[b+1] {
				table[b], table[b+1] = table[b+1], table[b]
			}
		}
	}
}
