package data

import (
	"github.com/go-redis/redis/v8"

	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/pkg/dbx/entx"
	"github.com/thoohv5/template/pkg/log"
)

type data struct {
	rdb *redis.Client
	edb *entx.ClientSet
}

// IData 数据源
type IData interface {
	// GetRdb redis
	GetRdb() *redis.Client
	// GetEdb db
	GetEdb() *entx.ClientSet
}

// New .
func New(c config.IConfig, log log.ILog) (IData, func(), error) {
	d := &data{
		// rdb: rdx.Open(c.GetRedis()),
	}
	// clientSet, err := dbx.OpenClientSet("main", c.GetDatabase())
	// if err != nil {
	// 	log.Error("db open err", zap.Error(err))
	//
	// 	return nil, nil, fmt.Errorf("db open err:%w", err)
	// }
	// d.edb = clientSet

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
func (d *data) GetEdb() *entx.ClientSet {
	return d.edb
}
