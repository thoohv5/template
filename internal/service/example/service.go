package example

import (
	"context"

	"github.com/thoohv5/template/internal/data/dao"
	"github.com/thoohv5/template/internal/data/dao/model"
	"github.com/thoohv5/template/pkg/dbx/standard"
)

type (
	service struct {
		log     standard.ILogger
		testDao dao.ITest
	}
	IService interface {
		Test(ctx context.Context, name string) error
	}
)

func New(log standard.ILogger, testDao dao.ITest) IService {
	return &service{
		log:     log,
		testDao: testDao,
	}
}

func (s *service) Test(ctx context.Context, name string) (err error) {
	return s.testDao.Add(ctx, &model.Test{
		Name: name,
	})
}
