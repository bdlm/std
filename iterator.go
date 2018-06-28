package std

/*
Iterator defines a simple interator interface.
*/
type Iterator interface {
	// Next reads the key and value at the current cursor position into pK
	// and pV respectively and moves the cursor forward one position. If
	// more data is available Next returns true, else false and resets the
	// cursor postion to the beginning of the data.
	Next(pK, pV *interface{}) bool

	// Reset sets the iterator position to 0.
	Reset()
}
