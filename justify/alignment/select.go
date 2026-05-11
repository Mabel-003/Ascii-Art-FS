package alignment

func SelectAlign(alignType string, ascii string, text string, banner string) string {
	switch alignType {
	case "right":
		return AlignRight(ascii)
	case "center":
		return AlignCenter(ascii)
	case "justify":
		return AlignJustify(text, banner)
	default:
		return AlignLeft(ascii)
	}
}