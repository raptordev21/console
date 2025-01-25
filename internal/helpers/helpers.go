package helpers

import "strings"

func GenerateSpaces(input string) string {
	length := len(input)
	return strings.Repeat(" ", length)
}
