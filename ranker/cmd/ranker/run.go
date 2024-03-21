package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	api_v1 "github.com/web3eye-io/Web3Eye/ranker/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Ranker daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		return logger.Init(logger.DebugLevel, config.GetConfig().Ranker.LogFile)
	},
	Action: func(c *cli.Context) error {
		logger.Sugar().Info("start launch")
		err := db.Init()
		if err != nil {
			panic(fmt.Errorf("mysql init err: %v", err))
		}
		logger.Sugar().Info("success to init db")

		err = milvusdb.Init(c.Context)
		if err != nil {
			panic(fmt.Errorf("milvus init err: %v", err))
		}
		logger.Sugar().Info("success to init milvus")

		go runGRPCServer(config.GetConfig().Ranker.GrpcPort)
		logger.Sugar().Info("success to start server")

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

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
