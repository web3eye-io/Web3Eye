package main

import (
	"os"
	"os/signal"
	"syscall"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/Web3Eye/config"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run dealer daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		return logger.Init(logger.DebugLevel, config.GetConfig().Dealer.LogFile)
	},
	Action: func(c *cli.Context) error {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		<-sigchan
		os.Exit(1)

		return nil
	},
}
