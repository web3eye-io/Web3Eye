package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/gateway/pkg/task"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Gateway daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		return logger.Init(logger.DebugLevel, config.GetConfig().Gateway.LogFile)
	},
	Action: func(c *cli.Context) error {

		config.GetConfig().CloudProxy.IP = "127.0.0.1"
		go task.Run(c.Context)

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		<-sigchan

		os.Exit(1)
		return nil
	},
}
