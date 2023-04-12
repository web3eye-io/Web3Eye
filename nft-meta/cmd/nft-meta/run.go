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
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/servermux"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	api_v1 "github.com/web3eye-io/Web3Eye/nft-meta/api/v1"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		return logger.Init(logger.DebugLevel, config.GetConfig().NFTMeta.LogFile)
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
		go runHTTPServer(config.GetConfig().NFTMeta.HTTPPort, config.GetConfig().NFTMeta.GrpcPort)
		go runGRPCServer(config.GetConfig().NFTMeta.GrpcPort)
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

func runHTTPServer(httpPort, grpcPort int) {
	httpEndpoint := fmt.Sprintf(":%v", httpPort)
	grpcEndpoint := fmt.Sprintf(":%v", grpcPort)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api_v1.RegisterGateway(mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service endpoint: %v", err)
	}

	servermux.AppServerMux().Handle("/v1/", mux)
	err = http.ListenAndServe(httpEndpoint, servermux.AppServerMux())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
