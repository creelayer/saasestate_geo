package helpers

import (
	"github.com/alexsergivan/transliterator"
)

type TransliteratorHelper struct {
	langCode string
	*transliterator.Transliterator
}

func NewTransliteratorHelper(langCode string) *TransliteratorHelper {
	c := &TransliteratorHelper{}
	c.Transliterator = transliterator.NewTransliterator(nil)
	c.langCode = langCode
	return c
}

func (c *TransliteratorHelper) Transliterate(text string) string {
	return c.Transliterator.Transliterate(text, c.langCode)
}
