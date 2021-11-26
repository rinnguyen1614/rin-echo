package middleware

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"rin-echo/common/domain"
	echox "rin-echo/common/echo"
	"rin-echo/common/log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thoas/go-funk"
)

type (
	RequestLoggerConfig struct {
		ApplicationName  string
		HeaderRequestID  string
		HeaderDeviceID   string
		HeaderDeviceName string
		// BeforeNextFunc defines a function that is called before next middleware or handler is called in chain.
		BeforeNextFunc func(c echo.Context)
		LogFunc        LogFunc
		Skipper        middleware.Skipper

		LogUserID             bool
		LogUserName           bool
		LogImpersonatorUserId bool
		LogOperation          bool
		LogRequestMethod      bool
		LogRequestURL         bool
		LogRequestID          bool
		LogRequestBody        bool
		LogLatency            bool
		LogLocation           bool
		LogIPAddress          bool
		LogDeviceID           bool
		LogDeviceName         bool
		LogBrowserInfo        bool
		LogResponseBody       bool
		LogStatusCode         bool
		LogError              bool
		LogRemark             bool
	}

	LogFunc func(c echo.Context, auditLog domain.AuditLog) error
)

var (
	DefaultRequestLogger = RequestLoggerConfig{
		Skipper:               middleware.DefaultSkipper,
		HeaderRequestID:       echox.HeaderRequestID,
		HeaderDeviceID:        echox.HeaderDeviceID,
		HeaderDeviceName:      echox.HeaderDeviceName,
		LogUserID:             true,
		LogUserName:           true,
		LogImpersonatorUserId: true,
		LogOperation:          true,
		LogRequestMethod:      true,
		LogRequestURL:         true,
		LogRequestID:          true,
		LogRequestBody:        true,
		LogLatency:            true,
		LogLocation:           true,
		LogIPAddress:          true,
		LogDeviceID:           true,
		LogDeviceName:         true,
		LogBrowserInfo:        true,
		LogStatusCode:         true,
		LogResponseBody:       false,
		LogError:              true,
		LogRemark:             true,
	}
)

func RequestLogger(applicationName string, logFunc LogFunc) echo.MiddlewareFunc {
	var config = DefaultRequestLogger
	config.ApplicationName = applicationName
	config.LogFunc = logFunc
	return RequestLoggerWithConfig(config)
}
func RequestLoggerWithOperation(applicationName string, operationName, operationMethod string, logFunc LogFunc) echo.MiddlewareFunc {
	var config = DefaultRequestLogger
	config.ApplicationName = applicationName
	config.LogFunc = logFunc
	return RequestLoggerWithConfig(config)
}

func RequestLoggerWithConfig(config RequestLoggerConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultRequestLogger.Skipper
	}

	if config.LogFunc == nil {
		panic("missing LogFunc callback function for request logger middleware")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			if config.BeforeNextFunc != nil {
				config.BeforeNextFunc(c)
			}

			var (
				cx, _              = echox.Contextx(c)
				logger             = c.Logger()
				now                = time.Now
				req                = c.Request()
				res                = c.Response()
				resBodyBuffer      = new(bytes.Buffer)
				userID             *uint
				username           string
				impersonatorUserId *uint
				operationName      string
				operationMethod    string
				requestMethod      string
				requestURL         string
				requestID          string
				requestBody        []byte
				startTime          = now()
				latency            time.Duration
				location           string
				ipAddress          string
				deviceID           string
				deviceName         string
				userAgent          string
				statusCode         int
				responseBody       []byte
				errorMsg           string
				remark             string
			)

			if config.LogResponseBody {
				mw := io.MultiWriter(res.Writer, resBodyBuffer)
				writer := &bodyResponseWriter{
					ResponseWriter: res.Writer,
					Writer:         mw,
				}
				res.Writer = writer
			}

			if err = next(c); err != nil {
				c.Error(err)
			}

			if (config.LogUserID || config.LogUserName) && cx != nil {
				if session, _ := cx.Session(); session != nil {
					if config.LogUserID {
						uID := session.UserID()
						userID = &uID
					}
					if config.LogUserName {
						if us := funk.Get(session, "Username"); us != nil {
							username = us.(string)
						}
					}
				}
			}

			if config.LogImpersonatorUserId {

			}
			if config.LogOperation && cx != nil {
				operationName, operationMethod = cx.Operation()
			}

			if config.LogRequestMethod {
				requestMethod = req.Method
			}
			if config.LogRequestURL {
				requestURL = req.RequestURI
			}
			if config.LogRequestID {
				requestID = req.Header.Get(config.HeaderRequestID)
				if requestID == "" {
					requestID = res.Header().Get(config.HeaderRequestID)
				}
			}
			if config.LogRequestBody {
				if req.Method != http.MethodGet && req.Body != nil {
					requestBody, err = ioutil.ReadAll(req.Body) // read
					if err != nil {
						echox.Logj(logger, log.ErrorLevel, map[string]interface{}{
							"message": "Read body from request error",
							"err":     err,
						})
					} else {
						req.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody)) // reset
					}
				}
			}
			if config.LogLatency {
				latency = now().Sub(startTime)
			}
			if config.LogLocation {

			}
			if config.LogIPAddress {
				ipAddress = c.RealIP()
			}
			if config.LogDeviceID {
				deviceID = req.Header.Get(config.HeaderDeviceID)
			}
			if config.LogDeviceName {
				deviceName = req.Header.Get(config.HeaderDeviceName)
			}
			if config.LogBrowserInfo {
				userAgent = req.UserAgent()
			}

			if resBodyBuffer != nil {
				responseBody = resBodyBuffer.Bytes()
			}

			if config.LogStatusCode {
				statusCode = res.Status

				if err != nil {
					if httpErr, ok := err.(*echo.HTTPError); ok {
						statusCode = httpErr.Code
					}
				}
			}
			if config.LogError && err != nil {
				errorMsg = err.Error()
			}
			if config.LogRemark {

			}

			v := domain.NewAuditLog(
				config.ApplicationName,
				userID,
				username,
				impersonatorUserId,
				operationName,
				operationMethod,
				requestMethod,
				requestURL,
				requestID,
				string(requestBody),
				startTime,
				latency,
				location,
				ipAddress,
				deviceID,
				deviceName,
				userAgent,
				string(responseBody),
				statusCode,
				errorMsg,
				remark,
			)

			if errOnLog := config.LogFunc(c, *v); errOnLog != nil {
				return errOnLog
			}

			return err
		}
	}
}

type bodyResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
