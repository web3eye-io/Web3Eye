//nolint
package main

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/cyber-tracer/config"
	"github.com/web3eye-io/cyber-tracer/utils"
)

func main() {
	logger.Init(logger.DebugLevel, "./")
	config.InitConfig()
	fmt.Println(utils.PrettyStruct(config.GetConfig().NFTMeta))
}
