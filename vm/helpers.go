package vm

type EncodedItem struct {
	Item      int
	LineCount int
}

/* run length encoding of line numbers. */
func RLEEncode(lines []EncodedItem, lineNumber int) []EncodedItem {
	if len(lines) == 0 || lines[len(lines)-1].Item != lineNumber {
		lines = append(lines, EncodedItem{Item: lineNumber, LineCount: 1})
	} else {
		lines[len(lines)-1].LineCount++
	}
	return lines
}
