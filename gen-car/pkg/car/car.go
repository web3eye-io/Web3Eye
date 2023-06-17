package car

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctfile"
	"github.com/web3eye-io/Web3Eye/common/unixfs"
)

const (
	DefaultCarVersion = 2
)

type CarInfo struct {
	FilePath    string
	RootCID     string
	CommPCID    string
	PayloadSize int64
	PieceSize   uint64
	Size        uint64
}

// CreateCar creates a car
// method from "github.com/ipld/go-car/cmd/car"
// the method is overwrite,because the original method is in main package
func CreateCar(ctx context.Context, carFilePath string, filesPath []string, version int) (carInfo *CarInfo, err error) {
	dir, _ := filepath.Split(carFilePath)
	tarFilePath := fmt.Sprintf("%v/%v.tar.gz", dir, uuid.NewString())
	err = ctfile.GenTarGZ(tarFilePath, filesPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := os.Remove(tarFilePath)
		logger.Sugar().Error(err)
	}()

	root, err := unixfs.CreateFilestore(context.Background(), tarFilePath, carFilePath)
	if err != nil {
		return nil, err
	}

	commp, err := ClientCalcCommP(ctx, carFilePath)
	if err != nil {
		return nil, err
	}

	return &CarInfo{
		FilePath:    carFilePath,
		RootCID:     root.String(),
		CommPCID:    commp.Root.String(),
		PayloadSize: commp.PayloadSize,
		PieceSize:   commp.PieceSize,
		Size:        commp.Size,
	}, err
}
