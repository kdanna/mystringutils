package mystringutils

import "strings"

//Upper returns the uppercase of the given string arg
func Upper(s string) string {
	return strings.ToUpper(s)
}

//Lower returns the lowercase of the given string arg
func Lower(s string) string {
	return strings.ToLower(s)
}
