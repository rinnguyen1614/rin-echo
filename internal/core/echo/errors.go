package echo

import core "github.com/rinnguyen1614/rin-echo/internal/core"

var (
	ERR_MISSING_CONTEXTX    = core.NewRinError("missing_contextx", "Missing Contextx! You should use middleware.Contextx")
	ERR_SESSION_NOT_FOUND   = core.NewRinError("session_not_found", "Session not found")
	ERR_LOCALIZER_NOT_FOUND = core.NewRinError("localizer", "Session not found")
)
