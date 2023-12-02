package text

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	c := cases.Title(language.English)
	return c.String(s)
}

func Lower(s string) string {
	c := cases.Lower(language.English)
	return c.String(s)
}

func Upper(s string) string {
	c := cases.Upper(language.English)
	return c.String(s)
}
