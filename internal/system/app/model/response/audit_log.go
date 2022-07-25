package response

import (
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/core/domain"
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
)

type AuditLog struct {
	model.CreationModel

	ApplicationName    string    `json:"application_name" `
	UserID             *uint     `json:"user_id" `
	Username           string    `json:"username" `
	ImpersonatorUserID *uint     `json:"impersonator_user_id" `
	OperationName      string    `json:"operation_name" `
	OperationMethod    string    `json:"operation_method" `
	RequestMethod      string    `json:"request_method" `
	RequestURL         string    `json:"request_url" `
	RequestID          string    `json:"request_id" `
	RequestBody        string    `json:"request_body"`
	StartTime          time.Time `json:"start_time" `
	Latency            int64     `json:"latency" `
	Location           string    `json:"location" `
	IPAddress          string    `json:"ip_address"`
	DeviceID           string    `json:"device_id" `
	DeviceName         string    `json:"device_name" `
	UserAgent          string    `json:"user_agent" `
	ResponseBody       string    `json:"response_body" `
	StatusCode         int       `json:"status_code" `
	Error              string    `json:"error"`
	Remark             string    `json:"remark" `
}

func NewAuditLog(e domain.AuditLog) AuditLog {
	return AuditLog{
		CreationModel:      model.NewCreationModelWithEntity(e.CreationEntity),
		ApplicationName:    e.ApplicationName,
		UserID:             e.UserID,
		Username:           e.Username,
		ImpersonatorUserID: e.ImpersonatorUserID,
		OperationName:      e.OperationName,
		OperationMethod:    e.OperationMethod,
		RequestMethod:      e.RequestMethod,
		RequestURL:         e.RequestURL,
		RequestID:          e.RequestID,
		RequestBody:        e.RequestBody,
		StartTime:          e.StartTime,
		Latency:            e.Latency,
		Location:           e.Location,
		IPAddress:          e.IPAddress,
		DeviceID:           e.DeviceID,
		DeviceName:         e.DeviceName,
		UserAgent:          e.UserAgent,
		ResponseBody:       e.ResponseBody,
		StatusCode:         e.StatusCode,
		Error:              e.Error,
		Remark:             e.Remark,
	}
}

type AuditLogs []*AuditLog
