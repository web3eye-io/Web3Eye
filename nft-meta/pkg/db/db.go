package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	// ent policy runtime
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/runtime"
)

const (
	maxLifeTime = time.Minute * 5
	maxConns    = 100
)

func client() (*ent.Client, error) {
	conn, err := GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

var (
	mu   = sync.Mutex{}
	conn *sql.DB
)

func GetConn() (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()
	if conn != nil {
		return conn, nil
	}
	var err error

	myConfig := config.GetConfig().MySQL

	withoutDBMSN := fmt.Sprintf("%v:%v@tcp(%v:%v)/?parseTime=true&interpolateParams=true",
		myConfig.User, myConfig.Password,
		myConfig.IP,
		myConfig.Port,
	)

	createSQL := fmt.Sprintf("create database if not exists %v;", myConfig.Database)
	conn, err := sql.Open("mysql", withoutDBMSN)
	if err != nil {
		logger.Sugar().Warnf("call Open error: %v", err)
		return nil, err
	}

	_, err = conn.Exec(createSQL)
	if err != nil {
		logger.Sugar().Warnf("exec sql failed: %v", err)
		return nil, err
	}
	conn.Close()

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		myConfig.User, myConfig.Password,
		myConfig.IP,
		myConfig.Port,
		myConfig.Database,
	)
	conn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Sugar().Warnf("call Open error: %v", err)
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.
	conn.SetConnMaxLifetime(maxLifeTime)
	conn.SetMaxOpenConns(maxConns)
	conn.SetMaxIdleConns(maxConns)

	return conn, nil
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}
	return cli.Schema.Create(context.Background())
}

func Client() (*ent.Client, error) {
	return client()
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return err
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return fmt.Errorf("fail get client transaction: %v", err)
	}

	succ := false
	defer func() {
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}
