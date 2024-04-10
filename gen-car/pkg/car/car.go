package car

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/filecoin-project/go-fil-markets/stores"
	"github.com/google/uuid"
	"github.com/ipfs/go-cidutil/cidenc"
	"github.com/ipld/go-car"
	selectorparse "github.com/ipld/go-ipld-prime/traversal/selector/parse"
	"github.com/multiformats/go-multibase"
	"github.com/web3eye-io/Web3Eye/common/ctfile"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/common/unixfs"
	"github.com/web3eye-io/Web3Eye/config"
)

const (
	DefaultCarVersion        = 2
	MaxTraversalLinks uint64 = 32 * (1 << 20)
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
	fileKey := fmt.Sprintf("%v.tar.gz", uuid.NewString())
	tarFilePath := fmt.Sprintf("%v/%v", dir, fileKey)
	err = ctfile.GenTarGZ(tarFilePath, filesPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := oss.UploadFile(ctx, tarFilePath, config.GetConfig().Minio.TarBucket, fileKey)
		if err != nil {
			logger.Sugar().Error(err)
		}

		err = os.Remove(tarFilePath)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}()

	tmp := fmt.Sprintf("%v/%v", config.GetConfig().GenCar.DataDir, uuid.NewString())
	root, err := unixfs.CreateFilestore(context.Background(), tarFilePath, tmp)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp)

	// open the positional reference CAR as a filestore.
	fs, err := stores.ReadOnlyFilestore(tmp)
	if err != nil {
		return nil, fmt.Errorf("failed to open filestore from carv2 in path %s: %w", tmp, err)
	}
	defer fs.Close()

	f, err := os.Create(carFilePath)
	if err != nil {
		return nil, err
	}

	// build a dense deterministic CAR (dense = containing filled leaves)
	if err := car.NewSelectiveCar(
		ctx,
		fs,
		[]car.Dag{{
			Root:     root,
			Selector: selectorparse.CommonSelector_ExploreAllRecursively,
		}},
		car.MaxTraversalLinks(MaxTraversalLinks),
	).Write(
		f,
	); err != nil {
		return nil, fmt.Errorf("failed to write CAR to output file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	encoder := cidenc.Encoder{Base: multibase.MustNewEncoder(multibase.Base32)}

	fmt.Println("Payload CID: ", encoder.Encode(root))

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
