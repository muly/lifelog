package lifelog

import (
	"strings"
)

// GetStringKey replaces the spaces in the given string with '-' inorder to prepare the string URL friendly
func StringKey(s string) string {
	// TODO: need to find what other characters are not url safe and eliminate them
	return strings.ToLower(strings.Replace(s, " ", "-", -1))
}
