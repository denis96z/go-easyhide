package parser

import (
	"strings"
)

func GetCommentText(s string) string {
	return strings.TrimSpace(
		strings.TrimPrefix(s, "//"),
	)
}
