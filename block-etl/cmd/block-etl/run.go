package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/eth"
	"github.com/web3eye-io/Web3Eye/config"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Block ETL daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		return logger.Init(logger.DebugLevel, config.GetConfig().BlockETL.LogFile)
	},
	Action: func(c *cli.Context) error {
		// TODO: choose block-etl type by environment varible
		indexer, err := eth.NewIndexer()
		if err != nil {
			return err
		}
		go indexer.IndexTasks(context.Background())

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		<-sigchan
		os.Exit(1)

		return nil
	},
}
