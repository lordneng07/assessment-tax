package helper

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CamelCase(s string) string {

	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	s = strings.ReplaceAll(s, "_", " ")

	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

	s = strings.ReplaceAll(s, " ", "")

	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}
