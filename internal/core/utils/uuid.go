package utils

import (
	"github.com/google/uuid"
)

type UUID = uuid.UUID

func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

func MustUUID() UUID {
	return uuid.New()
}

func MustString() string {
	return uuid.NewString()
}

func ParseUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}

func MustParseUUID(s string) UUID {
	uuid, err := ParseUUID(s)
	if err != nil {
		panic("uuid isn't valid")
	}
	return uuid
}

func ParseBytesUUID(b []byte) (UUID, error) {
	return uuid.ParseBytes(b)
}

func MustParseBytesUUID(b []byte) UUID {
	uuid, err := ParseBytesUUID(b)
	if err != nil {
		panic("uuid isn't valid")
	}
	return uuid
}
