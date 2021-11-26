package errors

import "rin-echo/common"

var (
	ErrRequestIDRequired = common.NewRinError("request_id_required", "request_id_required")
	ErrRequestIDInvalid  = common.NewRinError("request_id_invalid", "request_id_invalid")
)
