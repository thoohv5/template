package data

import (
	"github.com/go-redis/redis/v8"

	dbs "github.com/thoohv5/template/pkg/dbx/standard"
)

// IData 数据源
type IData interface {
	// GetRdb redis
	GetRdb() *redis.Client
	// GetEdb db
	GetEdb() dbs.IBuilder
}
