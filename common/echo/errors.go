package echo

import "rin-echo/common"

var (
	ERR_MISSING_CONTEXTX    = common.NewRinError("missing_contextx", "Missing Contextx! You should use middleware.Contextx")
	ERR_SESSION_NOT_FOUND   = common.NewRinError("session_not_found", "Session not found")
	ERR_LOCALIZER_NOT_FOUND = common.NewRinError("localizer", "Session not found")
)
