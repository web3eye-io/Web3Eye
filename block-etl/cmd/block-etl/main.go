package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/version"
	banner "github.com/common-nighthawk/go-figure"
	cli "github.com/urfave/cli/v2"
	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/servicename"
)

const (
	serviceName = "BlockETL"
	usageText   = "Block ETL Service"
)

func main() {
	commands := cli.Commands{runCmd}

	description := fmt.Sprintf(
		"%v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName,
		serviceName,
	)
	banner.NewColorFigure(serviceName, "", "green", true).Print()
	vesion, err := version.GetVersion()
	if err != nil {
		log.Fatalf("fail to get version, %v", err)
	}

	app := &cli.App{
		Name:        serviceName,
		Version:     vesion,
		Description: description,
		Usage:       usageText,
		Commands:    commands,
	}

	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", servicename.ServiceName, err)
		return
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("fail to run %v: %v", serviceName, err)
	}
}
