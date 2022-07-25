package query

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func ParseToken(str string) ([]Token, error) {
	var (
		reader      = strings.NewReader(str)
		token       Token
		tokens      []Token
		err         error
		openBrakets int
	)

	tokens = make([]Token, 0)
	for reader.Len() > 0 {
		token, err = readToken(reader)
		if err != nil {
			return nil, err
		}

		if token.Kind == UNKNOWN {
			continue
		}

		if token.Kind == CLAUSE {
			openBrakets++
		} else if token.Kind == CLAUSE_CLOSE {
			openBrakets--
		}
		tokens = append(tokens, token)
	}
	if openBrakets != 0 {
		return nil, errors.New("Unbalanced parenthesis")
	}

	return tokens, nil
}

func readToken(reader *strings.Reader) (token Token, err error) {
	var (
		tokenValue  interface{}
		tokenString string
		character   rune
		kind        TokenKind
	)

	// numeric is 0-9, or . or 0x followed by digits
	// string starts with '
	// variable is alphanumeric, always starts with a letter
	// bracket always means variable
	// symbols are anything non-alphanumeric

	for reader.Len() > 0 {
		character, _, _ = reader.ReadRune()

		if unicode.IsSpace(character) {
			continue
		}

		var negative = 1
		if character == '-' {
			character, _, _ = reader.ReadRune()
			negative = -1
		}

		if isNumeric(character) {
			if character == '0' && reader.Len() > 0 {
				character, _, _ = reader.ReadRune()
				if character == 'x' {
					tokenString, _ = readUntil(reader, false, true, isHexDigit)
					num, err := strconv.ParseUint(tokenString, 16, 64)
					if err != nil {
						return Token{}, errors.New(fmt.Sprintf("Unable to parse hex value '%v' to uint64", tokenString))
					}
					kind = UINT
					tokenValue = num * uint64(negative)

					break
				} else {
					reader.UnreadRune()
				}
			}

			tokenString, _ = readUntil(reader, false, true, isNumeric)
			iNum, err := strconv.ParseInt(tokenString, 10, 64)
			if err == nil {
				kind = INT
				tokenValue = iNum * int64(negative)
				break
			}

			fNum, err := strconv.ParseFloat(tokenString, 64)
			if err != nil {
				return Token{}, errors.New(fmt.Sprintf("Unable to parse numeric value '%v' to int64 or float64", tokenString))
			}

			kind = FLOAT
			tokenValue = fNum * float64(negative)
			break
		}

		if unicode.IsLetter(character) {
			tokenString, _ = readUntil(reader, false, true, isFieldName)

			switch v := strings.ToLower(tokenString); v {
			case "true":
				kind = BOOLEAN
				tokenValue = true
			case "false":
				kind = BOOLEAN
				tokenValue = false
			case "and", "or", "not":
				kind = LOGICAL_OPERATOR
				tokenValue = v
			case "in", "like":
				kind = CONDITION_OPERATOR
				tokenValue = v
			default:
				kind = FIELD
				tokenValue = tokenString
			}

			break
		}

		if !isNotQuote(character) {
			//move reader to next character (passed opened quote)
			reader.ReadRune()
			tokenString, completed := readUntil(reader, true, false, isNotQuote)
			if !completed {
				return Token{}, errors.New("Not found unclosed string literal")
			}

			// move reader to next character (passed closed quote)
			reader.ReadRune()

			tokenTime, err := parseTime(tokenString)
			if err == nil {
				kind = TIME
				tokenValue = tokenTime
			} else {
				kind = STRING
				tokenValue = tokenString
			}

			break
		}

		if character == ',' {
			tokenValue = string(character)
			kind = SEPARATOR
			break
		}

		if character == '(' {
			tokenValue = string(character)
			kind = CLAUSE
			break
		}

		if character == ')' {
			tokenValue = string(character)
			kind = CLAUSE_CLOSE
			break
		}

		tokenString, _ = readUntil(reader, false, true, isNotAlphanumeric)

		if _, ok := MapConditionOperators[tokenString]; ok {
			tokenValue = tokenString
			kind = CONDITION_OPERATOR
			break
		}

		if _, ok := MapLogicalOperators[tokenString]; ok {
			tokenValue = tokenString
			kind = LOGICAL_OPERATOR
			break
		}

		return Token{}, errors.New(fmt.Sprintf("Invalid token: '%s'", tokenString))
	}

	return Token{
		Kind:  kind,
		Value: tokenValue,
	}, nil
}

func readUntil(reader *strings.Reader, includeWhitespace, breakWhitespace bool, funcCondition func(rune) bool) (
	string, bool) {

	var (
		buff      bytes.Buffer
		character rune
		condition bool = false
	)

	// Unread rune with Seek instead
	// Because the reader.UnreadRune is useful the previous step which is the reader.ReadRune
	reader.Seek(-1, io.SeekCurrent)

	for reader.Len() > 0 {
		character, _, _ = reader.ReadRune()

		if unicode.IsSpace(character) {
			if breakWhitespace && buff.Len() > 0 {
				condition = true
				break
			}

			if !includeWhitespace {
				continue
			}
		}

		if funcCondition(character) {
			buff.WriteString(string(character))
		} else {
			condition = true // read all
			reader.UnreadRune()
			break
		}
	}

	return buff.String(), condition
}

func isHexDigit(r rune) bool {
	return unicode.IsDigit(r) ||
		r == 'a' ||
		r == 'b' ||
		r == 'c' ||
		r == 'd' ||
		r == 'e' ||
		r == 'f'
}

func isNumeric(r rune) bool {
	return unicode.IsDigit(r) || r == '.'
}

func isNotQuote(r rune) bool {
	return r != '\'' && r != '"'
}

func isNotAlphanumeric(r rune) bool {
	return !(unicode.IsDigit(r) ||
		unicode.IsLetter(r) ||
		r == '(' ||
		r == ')' ||
		!isNotQuote(r))
}

func isFieldName(r rune) bool {
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == '_' ||
		r == '.'
}

func parseTime(candidate string) (time.Time, error) {

	timeFormats := [...]string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",                         // RFC 3339
		"2006-01-02 15:04",                   // RFC 3339 with minutes
		"2006-01-02 15:04:05",                // RFC 3339 with seconds
		"2006-01-02 15:04:05-07:00",          // RFC 3339 with seconds and timezone
		"2006-01-02T15Z0700",                 // ISO8601 with hour
		"2006-01-02T15:04Z0700",              // ISO8601 with minutes
		"2006-01-02T15:04:05Z0700",           // ISO8601 with seconds
		"2006-01-02T15:04:05.999999999Z0700", // ISO8601 with nanoseconds
	}

	for _, format := range timeFormats {
		t, err := time.Parse(candidate, format)
		if err == nil {
			return t, nil
		}
	}

	return time.Now(), errors.New("Not found time's format")
}
