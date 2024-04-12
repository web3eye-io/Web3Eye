package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/common/servermux"
	"github.com/web3eye-io/Web3Eye/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	api_v1 "github.com/web3eye-io/Web3Eye/transform/api/v1"
	"github.com/web3eye-io/Web3Eye/transform/pkg/autototensor"
	"github.com/web3eye-io/Web3Eye/transform/pkg/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Transform daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		err := oss.Init(config.GetConfig().Minio.Region)
		if err != nil {
			return err
		}
		return logger.Init(logger.DebugLevel, config.GetConfig().Transform.LogFile)
	},
	Action: func(c *cli.Context) error {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		ctx, cancel := context.WithCancel(c.Context)
		defer cancel()
		go model.Run()
		go runGRPCServer(config.GetConfig().Transform.GrpcPort)
		go runHTTPServer(config.GetConfig().Transform.HTTPPort)
		go autototensor.Run(ctx)

		<-sigchan
		os.Exit(1)

		return nil
	},
}

func runGRPCServer(grpcPort int) {
	endpoint := fmt.Sprintf(":%v", grpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api_v1.Register(server)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runHTTPServer(httpPort int) {
	httpEndpoint := fmt.Sprintf(":%v", httpPort)
	mux := runtime.NewServeMux()
	servermux.AppServerMux().Handle("/v1/", mux)

	err := http.ListenAndServe(httpEndpoint, servermux.AppServerMux())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
