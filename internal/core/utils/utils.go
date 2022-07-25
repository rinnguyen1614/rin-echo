package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"reflect"
	"runtime"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/crypto/pbkdf2"
)

func DefaultValue(v interface{}, defaultValue interface{}) interface{} {
	rv := reflect.ValueOf(v)
	rDefaultValue := reflect.ValueOf(defaultValue)
	if rv.IsZero() {
		return defaultValue
	}

	if rv.Kind() == rDefaultValue.Kind() {
		return v
	}

	if rDefaultValue.Kind() == reflect.Ptr {
		ptr := reflect.New(rDefaultValue.Elem().Type())
		temp := ptr.Elem()
		temp.Set(rv)
		return ptr.Interface()
	}

	if rv.Kind() == reflect.Ptr {
		return rv.Elem().Interface()
	}

	return rv.Interface()
}

func Translate(localizer *i18n.Localizer, msgID, defaultMsg string) string {
	if localizer == nil {
		return defaultMsg
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: msgID,
		DefaultMessage: &i18n.Message{
			ID:    msgID,
			Other: defaultMsg,
		},
	})

	if err != nil {
		return err.Error()
	}

	return msg
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Encrypt(passphrase, plaintext string) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "_" + hex.EncodeToString(iv) + "_" + hex.EncodeToString(data)
}

func Decrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "_")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}
