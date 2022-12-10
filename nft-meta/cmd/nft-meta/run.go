package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/cyber-tracer/config"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/milvusdb"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/servermux"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/cyber-tracer/nft-meta/api"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db"
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
		err := logger.Init(logger.DebugLevel, config.GetConfig().NFTMeta.LogDir)
		if err != nil {
			panic(err)
		}
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
		err := db.Init()
		if err != nil {
			panic(fmt.Errorf("mysql init err: %v", err))
		}

		err = milvusdb.Init(context.Background())
		if err != nil {
			panic(fmt.Errorf("milvus init err: %v", err))
		}
		go runHTTPServer(config.GetConfig().NFTMeta.HTTPPort)
		runGRPCServer(config.GetConfig().NFTMeta.GrpcPort)
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
	api.Register(server)
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
