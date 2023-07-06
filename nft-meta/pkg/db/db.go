package db

import (
	"context"
	"database/sql"
	"fmt"
	"net"
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
	maxLifeTime     = time.Minute
	maxConns        = 100
	maxFailedConter = 3
)

var (
	mu            = sync.Mutex{}
	failedConter  = 0
	masterMysqlIP string
)

func client() (*ent.Client, error) {
	var err error
	if failedConter >= maxFailedConter {
		failedConter = 0
		masterMysqlIP, err = GetMasterIP()
		if err != nil {
			panic(err)
		}
	}

	conn, err := GetConn()
	if err != nil {
		failedConter++
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func GetConn() (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	myConfig := config.GetConfig().MySQL

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		myConfig.User, myConfig.Password,
		masterMysqlIP,
		myConfig.Port,
		myConfig.Database,
	)
	conn, err := sql.Open("mysql", dataSourceName)
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

func InitDatabase() error {
	mu.Lock()
	defer mu.Unlock()
	var err error

	myConfig := config.GetConfig().MySQL

	withoutDBMSN := fmt.Sprintf("%v:%v@tcp(%v:%v)/?parseTime=true&interpolateParams=true",
		myConfig.User, myConfig.Password,
		masterMysqlIP,
		myConfig.Port,
	)

	createSQL := fmt.Sprintf("create database if not exists %v;", myConfig.Database)
	conn, err := sql.Open("mysql", withoutDBMSN)
	if err != nil {
		logger.Sugar().Warnf("call Open error: %v", err)
		return err
	}

	_, err = conn.Exec(createSQL)
	if err != nil {
		logger.Sugar().Warnf("exec sql failed: %v", err)
		return err
	}
	return conn.Close()
}

func GetMasterIP() (string, error) {
	myConfig := config.GetConfig().MySQL
	ips, err := net.LookupHost(myConfig.Domain)
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		withoutDBMSN := fmt.Sprintf("%v:%v@tcp(%v:%v)/?parseTime=true&interpolateParams=true",
			myConfig.User, myConfig.Password,
			ip,
			myConfig.Port,
		)

		checkReadOnly := "SELECT @@read_only;"
		s := ""
		func() {
			conn, err := sql.Open("mysql", withoutDBMSN)
			if err != nil {
				logger.Sugar().Warnf("call Open error: %v, ip: %v", err, ip)
			}
			defer conn.Close()

			result := conn.QueryRow(checkReadOnly)
			if err != nil {
				logger.Sugar().Warnf("check read only failed: %v, ip: %v", err, ip)
			}
			err = result.Scan(&s)
			if err != nil {
				logger.Sugar().Warnf("check read only failed: %v, ip: %v", err, ip)
			}
		}()

		if s == "0" {
			return ip, nil
		}
	}
	return "", fmt.Errorf("cannot find mysql master node")
}

func Init() error {
	var err error
	masterMysqlIP, err = GetMasterIP()
	if err != nil {
		panic(err)
	}

	err = InitDatabase()
	if err != nil {
		panic(err)
	}

	cli, err := client()
	if err != nil {
		panic(err)
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
	defer cli.Close()

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
	defer cli.Close()

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}
