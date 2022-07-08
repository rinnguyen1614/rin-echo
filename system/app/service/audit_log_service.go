package service

import (
	"rin-echo/common/domain"
	echox "rin-echo/common/echo"
	"rin-echo/common/model"
	"rin-echo/common/query"
	"rin-echo/common/setting"
	uow "rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/app/model/response"

	"go.uber.org/zap"
)

type (
	AuditLogService interface {
		WithContext(echox.Context) AuditLogService

		Query(q *query.Query) (*model.QueryResult, error)

		Get(id uint) (response.AuditLog, error)
	}

	auditLogService struct {
		*echox.Service

		repo iuow.RepositoryOfEntity
	}
)

func NewAuditLogService(ux iuow.UnitOfWork, settingProvider setting.Provider, logger *zap.Logger) AuditLogService {
	return &auditLogService{
		Service: echox.NewService(ux, settingProvider, logger),

		repo: uow.NewRepositoryOfEntity(ux.DB(), &domain.AuditLog{}),
	}
}

func (s *auditLogService) WithContext(ctx echox.Context) AuditLogService {
	return &auditLogService{
		Service: s.Service.WithContext(ctx),
		repo:    s.repo,
	}
}

func (s auditLogService) Get(id uint) (response.AuditLog, error) {
	var auditLog domain.AuditLog
	if err := s.repo.GetID(&auditLog, id, nil); err != nil {
		return response.AuditLog{}, err
	}
	return response.NewAuditLog(auditLog), nil
}

func (s auditLogService) Query(q *query.Query) (*model.QueryResult, error) {
	var (
		queryBuilder, _ = uow.NewQueryBuilder(&domain.AuditLog{})
	)

	return q.QueryResult(s.repo, queryBuilder, nil, domain.AuditLog{}, response.AuditLog{})
}
