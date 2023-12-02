package text

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	return cases.Title(language.English).String(s)
}

func Lower(s string) string {
	return cases.Lower(language.English).String(s)
}

func Upper(s string) string {
	return cases.Upper(language.English).String(s)
}
