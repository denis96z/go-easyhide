package easyhide

import (
	"regexp"
)

type RegexpReplacement struct {
	Regexp *regexp.Regexp

	Replacement string
}
