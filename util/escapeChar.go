package util

import "strings"

func EscapeChar(term string) string {
	// Escape special characters for Redisearch query syntax
	escapedTerm := strings.ReplaceAll(term, ":", "\\:")       // Escape colon
	escapedTerm = strings.ReplaceAll(escapedTerm, " ", "\\ ") // Escape space
	escapedTerm = strings.ReplaceAll(escapedTerm, "@", "\\@") // Escape @
	escapedTerm = strings.ReplaceAll(escapedTerm, ".", "\\.") // Escape .
	// You can add more replacements as needed for other special characters
	return escapedTerm
}
