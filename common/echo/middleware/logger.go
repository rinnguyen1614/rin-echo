package middleware

import (
	"fmt"
	"time"

	echox "rin-echo/common/echo"
	logx "rin-echo/common/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	LoggerConfig struct {
		Skipper middleware.Skipper
	}
)

var (
	// DefaultLoggerConfig is the default Logger middleware config.
	DefaultLoggerConfig = LoggerConfig{
		Skipper: middleware.DefaultSkipper,
	}
)

func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			var (
				logger = c.Logger()
				res    = c.Response()
				start  = time.Now()
				fields = map[string]interface{}{}
			)

			if err = next(c); err != nil {
				fields["error"] = err
				c.Error(err)
			}

			generalFields(c, start, fields)

			if err != nil {
				if ll, ok := err.(logx.HasLogLevel); ok {
					echox.Logj(logger, ll.Level(), fields)
					return nil
				}
			}

			n := res.Status
			switch {
			case n >= 500:
				logger.Errorj(fields)
			case n >= 400:
				logger.Warnj(fields)
			case n >= 300:
				logger.Infoj(fields)
			default:
				logger.Infoj(fields)
			}
			return nil
		}
	}
}

func generalFields(c echo.Context, start time.Time, fields map[string]interface{}) {
	var (
		req             = c.Request()
		res             = c.Response()
		latency         = time.Now().Sub(start)
		requestID       = res.Header().Get(echox.HeaderRequestID)
		clientRequestID = req.Header.Get(echox.HeaderClientRequestID)
	)

	fields["remote_ip"] = c.RealIP()
	fields["start_time"] = start
	fields["user_agent"] = req.UserAgent()
	fields["host"] = req.Host
	fields["request"] = fmt.Sprintf("%s %s", req.Method, req.RequestURI)
	fields["status"] = res.Status
	fields["size"] = res.Size
	fields["latency"] = latency
	fields["latency_human"] = latency.String()
	// request id from server
	fields["request_id"] = requestID
	// request id from client
	if clientRequestID == "" {
		clientRequestID = res.Header().Get(echox.HeaderClientRequestID)
	}
	fields["client_request_id"] = clientRequestID
	if cc, _ := echox.Contextx(c); cc != nil {
		if session, _ := cc.Session(); session != nil {
			fields["user_id"] = session.UserID()
		}
	}
}
