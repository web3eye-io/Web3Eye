package main

import (
	"context"
	"fmt"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"

	v1 "github.com/web3eye-io/Web3Eye/gen-car/pkg/client/v1"
)

func main() {
	fmt.Println(v1.ReportFile(context.Background(), &npool.ReportFileRequest{
		ID:    "9d372318-339d-4d0a-9689-ef289ca34b2e",
		S3Key: "47522f9c-435f-493a-92e2-1b16cede8e7e.png",
	}))
}
