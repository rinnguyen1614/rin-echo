package en

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func RegisterDefaultTranslation(bundle *i18n.Bundle) {
	bundle.AddMessages(language.English, messages...)
}

var (
	messages = []*i18n.Message{
		{
			ID:    "required",
			Other: "{{ . | Field }} is a required field",
		},
		// len
		{
			ID:    "len_string",
			One:   "{{ . | Field }} must be in {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be in {{ . | Param}} characters in length",
		},
		{
			ID:    "len_items",
			One:   "{{ . | Field }} must contain {{ . | Param}} item",
			Other: "{{ . | Field }} must contain {{ . | Param}} items",
		},
		{
			ID:    "len_number",
			Other: "{{ . | Field }} must be equal {{ . | Param}} or greater",
		},
		// min
		{
			ID:    "min_string",
			One:   "{{ . | Field }} must be at least {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be at least {{ . | Param}} characters in length",
		},
		{
			ID:    "min_items",
			One:   "{{ . | Field }} must contain at least {{ . | Param}} item",
			Other: "{{ . | Field }} must contain at least {{ . | Param}} items",
		},
		{
			ID:    "min_number",
			Other: "{{ . | Field }} must be {{ . | Param}} or greater",
		},
		// max
		{
			ID:    "max_string",
			One:   "{{ . | Field }} must be a maximum of {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be a maximum of {{ . | Param}} characters in length",
		},
		{
			ID:    "max_items",
			One:   "{{ . | Field }} must contain at maximum {{ . | Param}} item",
			Other: "{{ . | Field }} must contain at maximum {{ . | Param}} items",
		},
		{
			ID:    "max_number",
			Other: "{{ . | Field }} must be equal {{ . | Param}} or less",
		},
		// eq
		{
			ID:    "eq",
			Other: "{{ . | Field }} is equal to {{ . | Param}}",
		},
		// ne
		{
			ID:    "ne",
			Other: "{{ . | Field }} should not be equal to {{ . | Param}}",
		},
		// lt
		{
			ID:    "lt_string",
			One:   "{{ . | Field }} must be less than {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be less than {{ . | Param}} characters in length",
		},
		{
			ID:    "lt_items",
			One:   "{{ . | Field }} must contain less than {{ . | Param}} item",
			Other: "{{ . | Field }} must contain less than {{ . | Param}} items",
		},
		{
			ID:    "lt_number",
			Other: "{{ . | Field }} must be less than {{ . | Param}}",
		},
		{
			ID:    "lt_datetime",
			Other: "{{ . | Field }} must be less than the current Date & Time",
		},
		// lte
		{
			ID:    "lte_string",
			One:   "{{ . | Field }} must be at maxium {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be at maxium {{ . | Param}} characters in length",
		},
		{
			ID:    "lte_items",
			One:   "{{ . | Field }} must contain at maxium {{ . | Param}} item",
			Other: "{{ . | Field }} must contain at maxium {{ . | Param}} items",
		},
		{
			ID:    "lte_number",
			Other: "{{ . | Field }} must be equal {{ . | Param}} or less",
		},
		{
			ID:    "lte_datetime",
			Other: "{{ . | Field }} must be less than or equal the current Date & Time",
		},
		// gt
		{
			ID:    "gt_string",
			One:   "{{ . | Field }} must be greater than {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be greater than {{ . | Param}} characters in length",
		},
		{
			ID:    "gt_items",
			One:   "{{ . | Field }} must contain more than {{ . | Param}} item",
			Other: "{{ . | Field }} must contain more than {{ . | Param}} items",
		},
		{
			ID:    "gt_number",
			Other: "{{ . | Field }} must be greater than {{ . | Param}}",
		},
		{
			ID:    "gt_datetime",
			Other: "{{ . | Field }} must be greater than the current Date & Time",
		},
		// gte
		{
			ID:    "gte_string",
			One:   "{{ . | Field }} must be at least {{ . | Param}} character in length",
			Other: "{{ . | Field }} must be at least {{ . | Param}} characters in length",
		},
		{
			ID:    "gte_items",
			One:   "{{ . | Field }} must contain at least {{ . | Param}} item",
			Other: "{{ . | Field }} must contain at least {{ . | Param}} items",
		},
		{
			ID:    "gte_number",
			Other: "{{ . | Field }} must be equal {{ . | Param}} or greater",
		},
		{
			ID:    "gte_datetime",
			Other: "{{ . | Field }} must be greater than or equal to the current Date & Time",
		},
		// eqfield
		{
			ID:    "eqfield",
			Other: "{{ . | Field }} must be equal to {{ . | Param}}",
		},
		// nefield
		{
			ID:    "nefield",
			Other: "{{ . | Field }} cannot be equal to {{ . | Param}}",
		},
		// gtfield
		{
			ID:    "gtfield",
			Other: "{{ . | Field }} must be greater than {{ . | Param}}",
		},
		// gtefield
		{
			ID:    "gtefield",
			Other: "{{ . | Field }} cmust be greater than or equal to {{ . | Param}}",
		},
		// ltfield
		{
			ID:    "ltfield",
			Other: "{{ . | Field }} must be less than {{ . | Param}}",
		},
		// ltefield
		{
			ID:    "ltefield",
			Other: "{{ . | Field }} cmust be less than or equal to {{ . | Param}}",
		},
		// eqcsfield
		{
			ID:    "eqcsfield",
			Other: "{{ . | Field }} must be equal to {{ . | Param}}",
		},
		// necsfield
		{
			ID:    "necsfield",
			Other: "{{ . | Field }} cannot be equal to {{ . | Param}}",
		},
		// gtcsfield
		{
			ID:    "gtcsfield",
			Other: "{{ . | Field }} must be greater than {{ . | Param}}",
		},
		// gtecsfield
		{
			ID:    "gtecsfield",
			Other: "{{ . | Field }} must be greater than or equal to {{ . | Param}}",
		},
		// ltcsfield
		{
			ID:    "ltcsfield",
			Other: "{{ . | Field }} must be less than {{ . | Param}}",
		},
		// ltecsfield
		{
			ID:    "ltecsfield",
			Other: "{{ . | Field }} must be less than or equal to {{ . | Param}}",
		},
		{
			ID:    "alpha",
			Other: "{{ . | Field }} can only contain alphabetic characters",
		},
		{
			ID:    "alphanum",
			Other: "{{ . | Field }} can only contain alphanumeric  characters",
		},
		{
			ID:    "numeric",
			Other: "{{ . | Field }} must be a valid numeric value",
		},
		{
			ID:    "number",
			Other: "{{ . | Field }} must be a valid number",
		},
		{
			ID:    "hexadecimal",
			Other: "{{ . | Field }} must be a valid hexadecimal",
		},
		{
			ID:    "hexcolor",
			Other: "{{ . | Field }} must be a valid HEX color",
		},
		{
			ID:    "rgb",
			Other: "{{ . | Field }} must be a valid RGB color",
		},
		{
			ID:    "rgba",
			Other: "{{ . | Field }} must be a valid RGBA color",
		},
		{
			ID:    "hsl",
			Other: "{{ . | Field }} must be a valid HSL color",
		},
		{
			ID:    "hsla",
			Other: "{{ . | Field }} must be a valid HSLA color",
		},
		{
			ID:    "e164",
			Other: "{{ . | Field }} must be a valid E.164 formatted phone number",
		},
		{
			ID:    "email",
			Other: "{{ . | Field }} must be a valid email address",
		},
		{
			ID:    "url",
			Other: "{{ . | Field }} must be a valid URL",
		},
		{
			ID:    "uri",
			Other: "{{ . | Field }} must be a valid URI",
		},
		{
			ID:    "base64",
			Other: "{{ . | Field }} must be a valid Base64 string",
		},
		{
			ID:    "contains",
			Other: "{{ . | Field }} must contain the text '{{ . | Param }}'",
		},
		{
			ID:    "containsany",
			Other: "{{ . | Field }} must contain at least one of the following characters '{{ . | Param }}'",
		},
		{
			ID:    "excludes",
			Other: "{{ . | Field }} cannot contain the text '{{ . | Param }}'",
		},
		{
			ID:    "excludesall",
			Other: "{{ . | Field }} cannot contain any of the following characters '{{ . | Param }}'",
		},
		{
			ID:    "excludesrune",
			Other: "{{ . | Field }} cannot contain the following '{{ . | Param }}'",
		},
	}
)
