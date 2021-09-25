package i18n

import (
	"rin-echo/common/validation/i18n/en"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func RegisterDefaultTranslation(buldle *i18n.Bundle, tag language.Tag) {
	switch tag {
	case language.English:
		en.RegisterDefaultTranslation(buldle)
	}
}
