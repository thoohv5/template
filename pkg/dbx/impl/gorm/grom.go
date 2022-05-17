package gorm

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/thoohv5/template/pkg/dbx/standard"
)

// gDb
type gDb struct {
	*standard.Options
	db *gorm.DB
}

func NewGDb() standard.IConnect {
	return &gDb{}
}

func CopyGDb(gdb *gorm.DB, sos ...standard.ServerOption) standard.IBuilder {
	opts := new(standard.Options)
	for _, so := range sos {
		so(opts)
	}

	if opts.GetIsSetLog() {
		gdb.Logger = WithLogger(gdb, opts.GetLogger())
	}

	return &gDb{
		Options: opts,
		db:      gdb,
	}
}

func (g *gDb) gDB() *gorm.DB {
	return g.db
}

func (g *gDb) Write() standard.IBuilder {
	// TODO implement me
	panic("implement me")
}

func (g *gDb) Read() standard.IBuilder {
	// TODO implement me
	panic("implement me")
}

type defaultLogger struct {
	logger Logger
	gdb    *gorm.DB
}

func New(zapLogger *zap.Logger) gormlogger.Interface {
	return &defaultLogger{
		logger: NewLogger(zapLogger),
	}
}

func WithLogger(gdb *gorm.DB, zapLogger *zap.Logger) gormlogger.Interface {
	return &defaultLogger{
		logger: NewLogger(zapLogger),
		gdb:    gdb,
	}
}

func (d *defaultLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return d.gdb.Logger.LogMode(level)
}

func (d *defaultLogger) Info(ctx context.Context, s string, i ...interface{}) {
	d.logger.Info(ctx, s, i...)
}

func (d *defaultLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	d.logger.Warn(ctx, s, i...)
}

func (d *defaultLogger) Error(ctx context.Context, s string, i ...interface{}) {
	d.logger.Error(ctx, s, i...)
}

func (d *defaultLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	d.logger.Trace(ctx, begin, fc, err)
}
