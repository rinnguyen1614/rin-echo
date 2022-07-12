package errors

import core "github.com/rinnguyen1614/rin-echo-core"

var (
	ErrRequestIDRequired = core.NewRinError("request_id_required", "request_id_required")
	ErrRequestIDInvalid  = core.NewRinError("request_id_invalid", "request_id_invalid")
)
