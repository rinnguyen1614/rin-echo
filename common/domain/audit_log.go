package domain

import (
	"time"
)

type AuditLog struct {
	CreationEntity

	ApplicationName    string    `gorm:"application_name;size:128"`
	UserID             *uint     `gorm:"column:user_id"`
	Username           string    `gorm:"column:username;size:255;"`
	ImpersonatorUserID *uint     `gorm:"column:impersonator_user_id "`
	OperationName      string    `gorm:"column:operation_name;size:128;"`
	OperationMethod    string    `gorm:"column:operation_method;size:128;"`
	RequestMethod      string    `gorm:"column:request_method;size:128;"`
	RequestURL         string    `gorm:"column:request_url;size:255;"`
	RequestID          string    `gorm:"column:request_id;size:255;"`
	RequestBody        string    `gorm:"column:request_body;type:text;"`
	StartTime          time.Time `gorm:"column:start_time;"`
	Latency            int64     `gorm:"column:latency;comment:Milliseconds"`
	Location           string    `gorm:"column:location;size:255;"`
	IPAddress          string    `gorm:"column:ip_address;size:128;"`
	DeviceID           string    `gorm:"column:device_id;size:128;"`
	DeviceName         string    `gorm:"column:device_name;size:128;"`
	UserAgent          string    `gorm:"column:user_agent;size:128;"`
	ResponseBody       string    `gorm:"column:response_body;type:text;"`
	StatusCode         int       `gorm:"column:status_code;size:4"`
	Error              string    `gorm:"column:error;size:255;"`
	Remark             string    `gorm:"column:remark;size:255;"`
}

func NewAuditLog(applicationName string,
	userID *uint,
	username string,
	impersonatorUserID *uint,
	operationName string,
	operationMethod string,
	requestMethod string,
	requestURL string,
	requestID string,
	requestBody string,
	startTime time.Time,
	latency time.Duration,
	location string,
	ipAddress string,
	deviceID string,
	deviceName string,
	userAgent string,
	responseBody string,
	statusCode int,
	errorMsg string,
	remark string) *AuditLog {
	return &AuditLog{
		ApplicationName:    applicationName,
		UserID:             userID,
		Username:           username,
		ImpersonatorUserID: impersonatorUserID,
		OperationName:      operationName,
		OperationMethod:    operationMethod,
		RequestMethod:      requestMethod,
		RequestURL:         requestURL,
		RequestID:          requestID,
		RequestBody:        requestBody,
		StartTime:          startTime,
		Latency:            latency.Milliseconds(),
		Location:           location,
		IPAddress:          ipAddress,
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		UserAgent:          userAgent,
		StatusCode:         statusCode,
		ResponseBody:       responseBody,
		Error:              errorMsg,
		Remark:             remark,
	}
}
