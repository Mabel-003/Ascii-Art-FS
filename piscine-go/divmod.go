package piscinego

func DivMod(a int, b int, div *int, mod *int) {
	division := a / b
	modulus := a % b
	*div = division
	*mod = modulus
}
