package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/servermux"
	"github.com/web3eye-io/cyber-tracer/config"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run NFT Meta daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "log dir",
			Aliases:     []string{"l"},
			Usage:       "log fir",
			EnvVars:     []string{"ENV_LOG_DIR"},
			Required:    false,
			Value:       "./",
			Destination: &logDir,
		},
	},
	Action: func(c *cli.Context) error {
		go runHTTPServer(config.GetConfig().BlockETL.HTTPPort)
		runGRPCServer(config.GetConfig().BlockETL.GrpcPort)
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
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runHTTPServer(httpPort int) {
	endpoint := fmt.Sprintf(":%v", httpPort)
	err := http.ListenAndServe(endpoint, servermux.AppServerMux())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
