package domain

import (
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/core/domain"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	"github.com/mssola/user_agent"
	"gorm.io/gorm"
)

type SecurityLog struct {
	domain.CreationEntity

	Username   string    `gorm:"column:username;size:255;"`
	Location   string    `gorm:"column:location;size:255;"`
	IPAddress  string    `gorm:"column:ip_address;size:128;"`
	DeviceID   string    `gorm:"column:device_id;size:128;"`
	DeviceName string    `gorm:"column:device_name;size:128;"`
	Browser    string    `gorm:"column:browser;size:128;"`
	Platform   string    `gorm:"column:platform;size:128;"`
	OS         string    `gorm:"column:os;size:128;"`
	UserAgent  string    `gorm:"column:user_agent;size:128;"`
	Time       time.Time `gorm:"column:time;"`
	StatusCode int       `gorm:"column:status_code;size:4"`
	Message    string    `gorm:"column:message;size:255;"`
}

func NewSecurityLog(username string,
	location string,
	ipAddress string,
	userAgent string,
	deviceID string,
	deviceName string,
	time time.Time,
	statusCode int,
	message string) *SecurityLog {
	var (
		agent                       = user_agent.New(userAgent)
		browserName, browserVersion = agent.Browser()
		browser                     = browserName + " " + browserVersion
		os                          = agent.OS()
		platform                    = agent.Platform()
		//model                       = agent.Model()
	)

	return &SecurityLog{
		Username:   username,
		Location:   location,
		IPAddress:  ipAddress,
		UserAgent:  userAgent,
		DeviceID:   deviceID,
		DeviceName: deviceName,
		Browser:    browser,
		Platform:   platform,
		OS:         os,
		Time:       time,
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewSecurityLogFromAuditLog(auditLog AuditLog) *SecurityLog {
	return NewSecurityLog(
		auditLog.Username,
		auditLog.Location,
		auditLog.IPAddress,
		auditLog.UserAgent,
		auditLog.DeviceID,
		auditLog.DeviceName,
		auditLog.StartTime,
		auditLog.StatusCode,
		"message",
	)
}

type SecurityLogRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) SecurityLogRepository
}
