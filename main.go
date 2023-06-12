package main

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	client "github.com/web3eye-io/Web3Eye/gen-car/pkg/client/v1"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
)

func main() {
	fmt.Println(oss.Init("default", config.GetConfig().Minio.TokenImageBucket))
	fmt.Println(client.ReportFile(context.Background(), &v1.ReportFileRequest{
		ID:    "0043793e-0129-4c89-a464-35d2b5f899c7",
		S3Key: "f1565d08-9153-46e9-9b86-2f954072785e.png",
	}))
}
