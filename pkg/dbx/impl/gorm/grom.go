package gorm

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	ds "github.com/thoohv5/template/pkg/dbx/standard"
)

// gDb
type gDb struct {
	*ds.Options
	db *gorm.DB
}

func NewGDb() ds.IConnect {
	return &gDb{}
}

func CopyGDb(gdb *gorm.DB, sos ...ds.ServerOption) ds.IBuilder {
	opts := new(ds.Options)
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

func (g *gDb) Write() ds.IBuilder {
	// TODO implement me
	panic("implement me")
}

func (g *gDb) Read() ds.IBuilder {
	// TODO implement me
	panic("implement me")
}

type defaultLogger struct {
	ds.ILogger
	Level                     gormlogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func WithLogger(gdb *gorm.DB, logger ds.ILogger) gormlogger.Interface {
	return &defaultLogger{
		ILogger:                   logger,
		SlowThreshold:             100 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
	}
}

func (d *defaultLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	d.Level = level
	return d
}

func (d *defaultLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if d.Level <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && d.Level >= gormlogger.Error && (!d.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		d.Errorc(ctx, "trace, err:%w, elapsed: %v, rows:%v, sql:%v", err, elapsed, rows, sql)
	case d.SlowThreshold != 0 && elapsed > d.SlowThreshold && d.Level >= gormlogger.Warn:
		sql, rows := fc()
		d.Warnc(ctx, "trace, err:%w, elapsed: %v, rows:%v, sql:%v", err, elapsed, rows, sql)
	case d.Level >= gormlogger.Info:
		sql, rows := fc()
		d.Debugc(ctx, "trace, err:%w, elapsed: %v, rows:%v, sql:%v", err, elapsed, rows, sql)
	}
}

func (d *defaultLogger) Info(ctx context.Context, s string, i ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Error(ctx context.Context, s string, i ...interface{}) {
	// TODO implement me
	panic("implement me")
}
