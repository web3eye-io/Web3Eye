package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/transform/pkg/autototensor"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	logger.Init(logger.InfoLevel, "./t.log")

	autototensor.Run()

	<-sigchan
	os.Exit(1)
}
