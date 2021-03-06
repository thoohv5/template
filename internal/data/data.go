package data

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	ds "github.com/thoohv5/template/internal/data/standard"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/dbx/adapter"
	dbs "github.com/thoohv5/template/pkg/dbx/standard"
	los "github.com/thoohv5/template/pkg/logx/standard"
)

type data struct {
	rdb *redis.Client
	edb dbs.IBuilder
}

// New .
func New(c config.IConfig, log los.ILogger) (ds.IData, func(), error) {
	d := &data{
		// rdb: rdx.Open(c.GetRedis()),
	}

	builder, err := adapter.GetConnect(adapter.Gorm).Connect(c.GetDatabase(), dbs.WithLogger(log))
	if err != nil {
		log.Error("db open err, err:%w", err)

		return nil, nil, fmt.Errorf("db open err:%w, config:%+v", err, c.GetDatabase())
	}
	d.edb = builder

	return d, func() {
		// if err := d.rdb.Close(); err != nil {
		// 	log.Error("redis close err", zap.Error(err))
		// 	panic(err)
		// }
		// if err := d.edb.Close(); err != nil {
		// 	log.Error("db close err", zap.Error(err))
		// 	panic(err)
		// }
	}, nil
}

// GetRdb redis
func (d *data) GetRdb() *redis.Client {
	return d.rdb
}

// GetEdb db
func (d *data) GetEdb() dbs.IBuilder {
	return d.edb
}
