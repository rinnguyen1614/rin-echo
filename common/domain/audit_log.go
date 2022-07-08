package domain

import (
	"time"
)

type AuditLog struct {
	CreationEntity

	ApplicationName    string    `json:"application_name" gorm:"application_name;size:128"`
	UserID             *uint     `json:"user_id" gorm:"column:user_id"`
	Username           string    `json:"username" gorm:"column:username;size:255;"`
	ImpersonatorUserID *uint     `json:"impersonator_user_id" gorm:"column:impersonator_user_id "`
	OperationName      string    `json:"operation_name" gorm:"column:operation_name;size:128;"`
	OperationMethod    string    `json:"operation_method" gorm:"column:operation_method;size:128;"`
	RequestMethod      string    `json:"request_method" gorm:"column:request_method;size:128;"`
	RequestURL         string    `json:"request_url" gorm:"column:request_url;size:255;"`
	RequestID          string    `json:"request_id" gorm:"column:request_id;size:255;"`
	RequestBody        string    `json:"request_body" gorm:"column:request_body;type:text;"`
	StartTime          time.Time `json:"start_time" gorm:"column:start_time;"`
	Latency            int64     `json:"latency" gorm:"column:latency;comment:Milliseconds"`
	Location           string    `json:"location" gorm:"column:location;size:255;"`
	IPAddress          string    `json:"ip_address" gorm:"column:ip_address;size:128;"`
	DeviceID           string    `json:"device_id" gorm:"column:device_id;size:128;"`
	DeviceName         string    `json:"device_name" gorm:"column:device_name;size:128;"`
	UserAgent          string    `json:"user_agent" gorm:"column:user_agent;size:128;"`
	ResponseBody       string    `json:"response_body" gorm:"column:response_body;type:text;"`
	StatusCode         int       `json:"status_code" gorm:"column:status_code;size:4"`
	Error              string    `json:"error" gorm:"column:error;size:255;"`
	Remark             string    `json:"remark" gorm:"column:remark;size:255;"`
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
