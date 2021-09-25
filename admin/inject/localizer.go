package inject

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func GetI18n() *i18n.Bundle {
	if service.i18n == nil {
		i18n := i18n.NewBundle(language.English)
		i18n.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

		for _, lang := range []string{"en"} {
			i18n.MustLoadMessageFile(fmt.Sprintf("resources/i18n/%v.yaml", lang))
		}

		service.i18n = i18n
	}

	return service.i18n
}
