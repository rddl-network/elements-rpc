package utils

import "strings"

func ParseResultToArray(result []byte) (resultArr []string) {
	resultStr := string(result)
	truncatedStr := resultStr[1 : len(resultStr)-1] // remove [] at beginning and end of Str
	sanitizedStr := strings.ReplaceAll(truncatedStr, "\"", "")

	return strings.Split(sanitizedStr, ",")
}
