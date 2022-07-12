package service

import (
	"rin-echo/system/app/model/response"

	"github.com/rinnguyen1614/rin-echo-core/domain"
	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/query"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	uow "github.com/rinnguyen1614/rin-echo-core/uow"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"

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
