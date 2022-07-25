package i18n

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/validation/i18n/en"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func RegisterDefaultTranslation(bundle *i18n.Bundle, tag language.Tag) {
	switch tag {
	case language.English:
		en.RegisterDefaultTranslation(bundle)
	}
}
