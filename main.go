package main

import (
	"context"
	"fmt"
	"time"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// contract.SetClientConnInterface(&contract.Po{})
	// fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
	conn, err := grpc.Dial("127.0.0.1:30121", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(1, err)
	}
	client := npool.NewManagerClient(conn)

	pc, err := client.ProxyChannel(context.TODO())
	if err != nil {
		fmt.Println(2, err)
	}

	for i := 0; i < 50; i++ {
		err = pc.Send(&npool.ProxyChannelRequest{
			MsgID: fmt.Sprint("ssssss", i),
		})
		if err != nil {
			fmt.Println(3, err)
		}
		time.Sleep(time.Second)
	}
}
