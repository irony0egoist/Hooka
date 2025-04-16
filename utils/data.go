package utils

// containsString is a helper function replicating slices.Contains for Go 1.20
// It checks if a string exists within a slice of strings.
func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true // Found the string
		}
	}
	return false // String not found after checking all items
}

// AppendString appends a string to a slice only if it's not already present.
// Compatible with Go 1.20.
func AppendString(slice []string, str string) []string {
	// Use the helper function instead of slices.Contains
	if !containsString(slice, str) {
		slice = append(slice, str)
	}
	return slice
}

// AppendSlice appends all strings from slice2 to slice1,
// skipping any strings that are already present in slice1.
// Compatible with Go 1.20.
func AppendSlice(slice, slice2 []string) []string {
	for _, entry := range slice2 {
		// Use the helper function instead of slices.Contains
		if !containsString(slice, entry) {
			slice = append(slice, entry)
		}
	}
	return slice
}
