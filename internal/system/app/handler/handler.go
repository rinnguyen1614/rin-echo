package handler

import (
	"strconv"

	"github.com/rinnguyen1614/rin-echo/internal/system/errors"
)

func CheckRequestIDParam(src string) (uint, error) {
	if len(src) == 0 {
		return 0, errors.ErrRequestIDRequired
	}
	id, err := strconv.Atoi(src)
	if err != nil {
		return 0, errors.ErrRequestIDInvalid
	}
	return uint(id), nil
}
