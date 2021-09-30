package dbx

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"
	
	"github.com/thoohv5/template/internal/ent"
	"github.com/thoohv5/template/internal/pkg/dbx/entx"
)

func fixMySQL(dsn string) (string, error) {
	c, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", err
	}

	// 通用参数
	c.ParseTime = true
	c.InterpolateParams = true
	c.Collation = "utf8mb4_general_ci"
	c.Loc = time.UTC
	delete(c.Params, "charset") // 不需要 charset, 设置好 collation 即可
	return c.FormatDSN(), nil
}

func OpenClient(resourceName string, cfg *DatabaseConfig) (*ent.Client, error) {
	var err error

	dsn := cfg.DSN
	driverName := cfg.Driver

	if strings.HasPrefix(driverName, "mysql") {
		dsn, err = fixMySQL(dsn)
		if err != nil {
			return nil, err
		}
	}

	driver := instrumentedsql.WrapDriver(&mysql.MySQLDriver{},
		instrumentedsql.WithTracer(opentracing.NewTracer(true)),
		instrumentedsql.WithOpsExcluded(
			instrumentedsql.OpSQLPing,
			instrumentedsql.OpSQLDummyPing,
			instrumentedsql.OpSQLConnectorConnect,
			instrumentedsql.OpSQLRowsNext,
		))
	wrappedDriverName := fmt.Sprintf("%s-otsql-%s", driverName, resourceName)
	sql.Register(wrappedDriverName, driver)

	db, err := sql.Open(wrappedDriverName, dsn)
	if err != nil {
		return nil, err
	}

	if cfg.MaxIdleConns != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxIdleTime != 0 {
		dbSetConnMaxIdleTime(db, time.Second*time.Duration(cfg.ConnMaxLifeTime))
	}
	if cfg.ConnMaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(cfg.ConnMaxLifeTime))
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(driverName, db)
	clientOpts := []ent.Option{ent.Driver(drv)}
	if cfg.Debug {
		clientOpts = append(clientOpts, ent.Driver(dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
			fmt.Printf("SQL--%v\n", i)
		})))
	}
	client := ent.NewClient(clientOpts...)
	return client, nil
}

func OpenClientSet(resourceName string, cfg *Configs) (*entx.ClientSet, error) {
	dbPrimary, err := OpenClient(resourceName, cfg.DatabaseConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open primary db: %w", err)
	}

	var replicas []*ent.Client
	for name, replicaCfg := range cfg.Replicas {
		replica, err := OpenClient(resourceName, replicaCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to open replica db %s: %w", name, err)
		}

		replicas = append(replicas, replica)
	}

	cs := entx.NewClientSet(dbPrimary, entx.WithReplicas(replicas))
	return cs, nil
}
