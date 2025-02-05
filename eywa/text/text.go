package text

import (
	"github.com/Lofanmi/pinyin-golang/pinyin"
	opencc "github.com/liuzl/gocc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(str string) string {
	return cases.Title(language.English).String(str)
}

func Lower(str string) string {
	return cases.Lower(language.English).String(str)
}

func Upper(str string) string {
	return cases.Upper(language.English).String(str)
}

func Pinyin(str string) string {
	dict := pinyin.NewDict()
	return dict.Sentence(str).Unicode()
}

func S2T(str string) string {
	s2t, err := opencc.New("s2t")
	if err != nil {
		return err.Error()
	}
	out, err := s2t.Convert(str)
	if err != nil {
		return err.Error()
	}

	return out
}
