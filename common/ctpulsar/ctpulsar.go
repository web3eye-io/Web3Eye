package ctpulsar

import (
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/config"
)

func Client() (pulsar.Client, error) {
	pulsarConfig := config.GetConfig().Pulsar
	pulsarUrl := fmt.Sprintf(
		"pulsar://%v:%v",
		pulsarConfig.Domain,
		pulsarConfig.Port,
	)

	return pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarUrl,
		OperationTimeout:  time.Duration(pulsarConfig.OperationTimeout) * time.Second,
		ConnectionTimeout: time.Duration(pulsarConfig.ConnectionTimeout) * time.Second,
	})
}
