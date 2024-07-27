package util

// Checks for an error and panics if one is present.
func PanicCheck(e error) {
	if e != nil {
		panic(e)
	}
}
