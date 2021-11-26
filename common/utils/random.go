package utils

import (
	"math/rand"
	"time"
)

const (
	symbol  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"
	letter  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	numeric = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(length int, s string) string {
	var chars = []byte(s)
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func RandomSymbol(length int) string {
	return RandomString(length, symbol)
}

func RandomLetter(length int) string {
	return RandomString(length, letter)
}

func RandomNumeric(length int) string {
	return RandomString(length, numeric)
}
